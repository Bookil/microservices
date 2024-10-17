package grpc

import (
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
	port int
	productv1.UnimplementedProductServiceServer
}

func NewAdapter(api ports.APIPort, auth ports.AuthPort,validator ports.Validation, port int) *Adapter {
	return &Adapter{
		api:             api,
		port:            port,
		validator: validator,
		authInterceptor: interceptor.NewAuthInterceptor(auth),
	}
}

func (a *Adapter) Run() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %d, error: %v", a.port, err)
	}

	grpcServer := grpc.NewServer(
		// grpc.UnaryInterceptor(a.authInterceptor.AuthInterceptor),
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
