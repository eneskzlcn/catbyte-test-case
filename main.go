package main

import (
	"github.com/eneskzlcn/catbyte-test-task/internal/message"
	"github.com/eneskzlcn/catbyte-test-task/internal/processor"
	"github.com/eneskzlcn/catbyte-test-task/rabbitmq"
	"github.com/eneskzlcn/catbyte-test-task/redis"
	server "github.com/eneskzlcn/catbyte-test-task/server"
	"log"
)

func main() {
	rabbitMQClient := rabbitmq.New("amqp://guest:guest@localhost:6001/", "catbyte-messages")
	redisClient := redis.NewClient("localhost:6379", "")

	messageService := message.NewService(rabbitMQClient)
	messageHandler := message.NewHandler(messageService)

	processorService := processor.NewService(rabbitMQClient, redisClient)

	go processorService.StartProcessing()

	server := server.NewServer([]server.Handler{
		messageHandler,
	}, []string{"*"}, "3000")

	err := server.Start()
	if err != nil {
		log.Println("server closed with error")
	}
}
