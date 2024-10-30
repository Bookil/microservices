package main

import (
	"log"

	"product/config"
	"product/internal/adapters/auth"
	"product/internal/adapters/cart"
	"product/internal/adapters/db"
	"product/internal/adapters/db/mysql_adapter"
	"product/internal/adapters/grpc"
	"product/internal/adapters/validation"
	"product/internal/application/core/api"
)

func main() {
	configs := config.Read()

	db, err := db.NewDB(&configs.Mysql)
	checkError(err)

	dbAdapter := mysql_adapter.NewAdapter(db)

	cartAdapter, err := cart.NewAdapter(&configs.CartService)
	checkError(err)

	authAdapter, err := auth.NewAdapter(&configs.AuthService)
	checkError(err)

	application := api.NewApplication(cartAdapter, dbAdapter)

	validator := validation.NewValidator()

	grpcAdapter := grpc.NewAdapter(application, authAdapter, validator, configs.Server.Port)

	grpcAdapter.Run()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
