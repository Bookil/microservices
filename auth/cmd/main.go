package main

import (
	"log"

	"github.com/Bookil/microservices/auth/config"
	"github.com/Bookil/microservices/auth/internal/adapters/db"
	"github.com/Bookil/microservices/auth/internal/adapters/db/mysql_adapter"
	"github.com/Bookil/microservices/auth/internal/adapters/grpc"
	"github.com/Bookil/microservices/auth/internal/application/core/api"
	"github.com/Bookil/microservices/auth/utils/hash"
	auth_manager "github.com/tahadostifam/go-auth-manager"
)

func main() {
	config := config.Read()

	mysqlAdapter, err := mysql_adapter.NewAdapter(&config.Mysql)
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	redisClient := db.GetRedisInstance(config.Redis)

	authManger := auth_manager.NewAuthManager(redisClient, auth_manager.AuthManagerOpts{
		PrivateKey: config.JWT.SecretKey,
	})

	hashManger := hash.NewHashManager(hash.DefaultHashParams)

	api := api.NewApplication(mysqlAdapter, authManger, hashManger)

	server := grpc.NewAdapter(api, config.Server.Port)

	server.Run()
}
