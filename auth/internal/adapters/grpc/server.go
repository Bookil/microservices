package grpc

import (
	"fmt"
	"log"
	"net"

	authv1 "github.com/Bookil/Bookil-Proto/gen/golang/auth/v1"

	"github.com/Bookil/microservices/auth/config"
	"github.com/Bookil/microservices/auth/internal/ports"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Adapter struct {
	api       ports.APIPort
	validator ports.Validation
	port      int
	authv1.UnimplementedAuthServiceServer
}

func NewAdapter(api ports.APIPort, validator ports.Validation, port int) *Adapter {
	return &Adapter{api: api, validator: validator, port: port}
}

func (a *Adapter) Run() {
	var err error

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %d, error: %v", a.port, err)
	}

	grpcServer := grpc.NewServer()
	authv1.RegisterAuthServiceServer(grpcServer, a)

	if config.CurrentEnv == config.Development {
		reflection.Register(grpcServer)
	}

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve grpc on port ")
	}
}
