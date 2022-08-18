package main

import (
	"github.com/eneskzlcn/catbyte-test-task/internal/config"
	"github.com/eneskzlcn/catbyte-test-task/internal/message"
	"github.com/eneskzlcn/catbyte-test-task/internal/processor"
	"github.com/eneskzlcn/catbyte-test-task/internal/reporting"
	"github.com/eneskzlcn/catbyte-test-task/rabbitmq"
	"github.com/eneskzlcn/catbyte-test-task/redis"
	"github.com/eneskzlcn/catbyte-test-task/server"
	"log"
	"os"
)

func main() {
	//does not in any other environment (prod, qa or etc.) yet. So directly load local config.
	config, err := config.LoadConfig(".dev/", "local", "yaml")
	if err != nil {
		return
	}

	rabbitMQClient := rabbitmq.New(config.RabbitMQ.Address, config.RabbitMQ.Queue)
	redisClient := redis.NewClient(config.Redis.Address, config.Redis.Password)

	//app 1: message api
	messageService := message.NewService(rabbitMQClient)
	messageHandler := message.NewHandler(messageService)

	//app 2: processor
	processorService := processor.NewService(rabbitMQClient, redisClient)
	go processorService.StartProcessing()

	//app 3: reporting
	reportingService := reporting.NewService(redisClient)
	reportingHandler := reporting.NewHandler(reportingService)

	server := server.NewServer([]server.Handler{
		messageHandler,
		reportingHandler,
	}, config.Server.TrustedProxies, config.Server.Port)

	if err = server.Start(); err != nil {
		log.Println("server closed with error")
		os.Exit(1)
	}
}
