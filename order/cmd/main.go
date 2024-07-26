package main

import (
	"log"

	"github.com/Bookil/microservices/order/config"
	"github.com/Bookil/microservices/order/internal/adapters/db"
	"github.com/Bookil/microservices/order/internal/adapters/grpc"
	"github.com/Bookil/microservices/order/internal/adapters/payment"
	"github.com/Bookil/microservices/order/internal/application/core/api"
)

func main() {
	config := config.Read() 

	dbAdapter, err := db.NewAdapter(&config.Mysql)
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	paymentAdapter, err := payment.NewAdapter(&config.PaymentService)
	if err != nil {
		log.Fatalf("Failed to initialize payment stub. Error: %v", err)
	}

	application := api.NewApplication(dbAdapter, paymentAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.Server.Port)

	grpcAdapter.Run()
}
