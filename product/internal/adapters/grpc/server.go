package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"product/config"
	"product/internal/adapters/grpc/interceptor"
	"product/internal/ports"

	productv1 "github.com/Bookil/Bookil-Proto/gen/golang/product/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Adapter struct {
	api             ports.APIPort
	authInterceptor *interceptor.AuthInterceptor
	validator       ports.Validation
	port            int
	productv1.UnimplementedProductServiceServer
}

func NewAdapter(api ports.APIPort, auth ports.AuthPort, validator ports.Validation, port int) *Adapter {
	return &Adapter{
		api:             api,
		port:            port,
		validator:       validator,
		authInterceptor: interceptor.NewAuthInterceptor(auth),
	}
}

func (a *Adapter) Run() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %d, error: %v", a.port, err)
	}

	combinedUnaryInterceptor := grpc.UnaryInterceptor(chainUnaryInterceptors(
		a.authInterceptor.AuthInterceptor,
		a.authInterceptor.RoleAuthInterceptor,
	))

	grpcServer := grpc.NewServer(
		combinedUnaryInterceptor,
	)

	productv1.RegisterProductServiceServer(grpcServer, a)

	if config.CurrentEnv == config.Development {
		reflection.Register(grpcServer)
	}

	log.Printf("Starting gRPC server on %d\n", a.port)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func chainUnaryInterceptors(interceptors ...grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		chain := handler
		for i := len(interceptors) - 1; i >= 0; i-- {
			interceptor := interceptors[i]
			chain = buildChain(interceptor, chain, info)
		}
		return chain(ctx, req)
	}
}

// buildChain builds a unary interceptor chain.
func buildChain(interceptor grpc.UnaryServerInterceptor, next grpc.UnaryHandler, info *grpc.UnaryServerInfo) grpc.UnaryHandler {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return interceptor(ctx, req, info, next)
	}
}
