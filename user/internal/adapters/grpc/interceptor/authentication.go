package interceptor

import (
	"context"

	"github.com/Bookil/microservices/user/internal/ports"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const accessTokenHeader = "X-Access-Token"

type UserID struct{}

type AuthInterceptor struct {
	auth ports.AuthPort
}

func NewAuthInterceptor(auth ports.AuthPort) *AuthInterceptor {
	return &AuthInterceptor{
		auth: auth,
	}
}

func (a *AuthInterceptor) AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// List of methods that require authentication
	authMethods := map[string]bool{
		"/userv1.UserService/Update":         true,
		"/userv1.UserService/ChangePassword": true,
	}

	if authMethods[info.FullMethod] {
		// Perform authentication logic here
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, "missing metadata")
		}

		if token, ok := md[accessTokenHeader]; ok && len(token) > 0 {
			userID, err := a.auth.Authenticate(ctx, token[0])
			if err != nil {
				return nil, status.Errorf(codes.Unauthenticated, err.Error())
			}
			if len(userID) > 0 {
				ctx = context.WithValue(ctx, UserID{}, userID)
				return handler(ctx, req)
			}
		}

		return nil, status.Errorf(codes.Unauthenticated, "access dined")
	}

	return handler(ctx, req)
}
