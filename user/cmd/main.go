package main

import (
	"log"

	"github.com/Bookil/microservices/user/config"
	"github.com/Bookil/microservices/user/internal/adapters/auth"
	"github.com/Bookil/microservices/user/internal/adapters/db"
	"github.com/Bookil/microservices/user/internal/adapters/db/mysql_adapter"
	"github.com/Bookil/microservices/user/internal/adapters/grpc"
	"github.com/Bookil/microservices/user/internal/adapters/validation"
	"github.com/Bookil/microservices/user/internal/application/core/api"
)

func main() {
	configs := config.Read()

	auth, err := auth.NewAdapter(&configs.Auth)
	if err != nil {
		log.Fatalf("error getting auth adapter:%v", err)
	}

	db, err := db.NewDB(&configs.Mysql)
	if err != nil {
		log.Fatalf("error getting DB instance:%v", err)
	}

	mysqlAdapter, err := mysql_adapter.NewAdapter(db)
	if err != nil {
		log.Fatalf("error getting db adapter:%v", err)
	}

	api := api.NewApplication(auth, mysqlAdapter)

	validator := validation.NewValidator()

	grpcAdapter := grpc.NewAdapter(api, auth, validator, configs.Server.Port)

	grpcAdapter.Run()
}
