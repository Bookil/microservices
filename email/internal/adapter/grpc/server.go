package grpc

import (
	"fmt"
	"log"
	"net"

	"email/config"
	"email/internal/ports"

	emailv1 "github.com/Bookil/Bookil-Proto/gen/golang/email/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Adapter struct {
	api  ports.APIPort
	port int
	emailv1.UnimplementedEmailServiceServer
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{
		api:  api,
		port: port,
	}
}

func (a *Adapter) Run() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %d, error: %v", a.port, err)
	}

	grpcServer := grpc.NewServer()

	emailv1.RegisterEmailServiceServer(grpcServer, a)

	if config.CurrentEnv == config.Development {
		reflection.Register(grpcServer)
	}

	log.Printf("Starting gRPC server on %d\n", a.port)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
