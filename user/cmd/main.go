package main

import (
	"log"

	"github.com/Bookil/microservices/user/config"
	"github.com/Bookil/microservices/user/internal/adapters/auth"
	"github.com/Bookil/microservices/user/internal/adapters/db/mysql_adapter"
	"github.com/Bookil/microservices/user/internal/adapters/email"
	"github.com/Bookil/microservices/user/internal/adapters/grpc"
	"github.com/Bookil/microservices/user/internal/application/core/api"
)

func main(){
	configs := config.Read()


	email := email.NewEmailAdapter()

	auth,err := auth.NewAdapter(&configs.Auth)
	if err != nil {
		log.Fatalf("error getting auth adapter:%v",err)
	}

	db,err := mysql_adapter.NewAdapter(&configs.Mysql)
	if err != nil {
		log.Fatalf("error getting db adapter:%v",err)
	}

	api := api.NewApplication(auth,email,db)

	grpcAdapter := grpc.NewAdapter(api,auth,configs.Server.Port)

	grpcAdapter.Run()
}