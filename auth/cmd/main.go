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
	"github.com/Bookil/microservices/auth/internal/ports"
)

func main() {
	configs := config.Read()

	mysqlAdapter, err := mysql_adapter.NewAdapter(&configs.Mysql)
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	redisClient := db.GetRedisInstance(configs.Redis)

	userService, err := user.NewAdapter(&configs.UserService)
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	var emailService ports.EmailPort

	// tmp
	isProduct := config.CurrentEnv == config.Production

	log.Println("env:", isProduct)

	if isProduct {
		emailProductService, err := email.NewAdapter(&configs.EmailService)
		if err != nil {
			log.Fatalf("Failed to connect to email service:%v", err)
		}

		emailService = emailProductService
	} else {
		emailService = email.NewDevEmailAdapter()
	}

	authManger := auth_manager.NewAdapter(redisClient, configs.JWT)

	hashManger := hash.NewHashManager(hash.DefaultHashParams)

	validator := validation.NewValidator()

	api := api.NewApplication(mysqlAdapter, userService, emailService, authManger, hashManger)

	server := grpc.NewAdapter(api, validator, configs.Server.Port)

	server.Run()
}
