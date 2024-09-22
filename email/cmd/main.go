package main

import (
	"email/config"
	"email/internal/adapter/grpc"
	"email/internal/adapter/smtp"
	"email/internal/application/core/api"
)

func main() {
	config := config.Read()

	smtp := smtp.NewSMTPAdapter(&config.SMTP)

	api := api.NewApplication(smtp)

	grpc := grpc.NewAdapter(api, config.Server.Port)

	grpc.Run()
}
