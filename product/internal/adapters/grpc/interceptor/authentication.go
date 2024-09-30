package interceptor

import (
	"context"

	"product/internal/ports"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const accessTokenHeader = "x-access-token"

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
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "missing metadata")
	}

	token, ok := md[accessTokenHeader]
	if !ok && len(token) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "access dined")
	}

	userID, err := a.auth.Authenticate(ctx, token[0])
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}

	if len(userID) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "access dined")
	}

	ctx = context.WithValue(ctx, UserID{}, userID)
	return handler(ctx, req)
}
