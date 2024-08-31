package main

import (
	"log"

	"github.com/Bookil/microservices/auth/config"
	"github.com/Bookil/microservices/auth/internal/adapters/auth_manager"
	"github.com/Bookil/microservices/auth/internal/adapters/db"
	"github.com/Bookil/microservices/auth/internal/adapters/db/mysql_adapter"
	"github.com/Bookil/microservices/auth/internal/adapters/email"
	"github.com/Bookil/microservices/auth/internal/adapters/grpc"
	"github.com/Bookil/microservices/auth/internal/adapters/hash"
	"github.com/Bookil/microservices/auth/internal/adapters/user"
	"github.com/Bookil/microservices/auth/internal/adapters/validation"
	"github.com/Bookil/microservices/auth/internal/application/core/api"
)

func main() {
	config := config.Read()

	mysqlAdapter, err := mysql_adapter.NewAdapter(&config.Mysql)
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	redisClient := db.GetRedisInstance(config.Redis)

	userService,err := user.NewAdapter(&config.UserService)
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	emailService := email.NewEmailAdapter()
	authManger := auth_manager.NewAdapter(redisClient,config.JWT)


	hashManger := hash.NewHashManager(hash.DefaultHashParams)

	validator := validation.NewValidator()

	api := api.NewApplication(mysqlAdapter, userService,emailService,authManger, hashManger)

	server := grpc.NewAdapter(api, validator, config.Server.Port)

	server.Run()
}
