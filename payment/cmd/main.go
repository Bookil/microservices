package main

import (
	"log"

	"github.com/Bookil/Bookil-Microservices/payment/config"
	"github.com/Bookil/Bookil-Microservices/payment/internal/adapter/db"
	"github.com/Bookil/Bookil-Microservices/payment/internal/adapter/grpc"
	"github.com/Bookil/Bookil-Microservices/payment/internal/application/core/api"
)

func main() {
	config := config.Read()

	dbAdapter, err := db.NewAdapter(&config.Mysql)
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	if err != nil {
		log.Fatalf("Failed to initialize payment stub. Error: %v", err)
	}

	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.Server.Port)

	grpcAdapter.Run()
}
