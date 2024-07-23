package cmd

import (
	"log"

	"github.com/Bookil/Bookil-Microservices/payment/config"
	"github.com/Bookil/Bookil-Microservices/payment/internal/adapter/db"
	"github.com/Bookil/Bookil-Microservices/payment/internal/adapter/grpc"
	"github.com/Bookil/Bookil-Microservices/payment/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application,config.GetApplicationPort())
	grpcAdapter.Run()
}
