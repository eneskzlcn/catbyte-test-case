package processor

import (
	"encoding/json"
	"fmt"
)

const Consumer = "processor"

type RabbitMQClient interface {
	Consume(onReceived chan []byte, consumer string)
}
type RedisClient interface {
	SaveToArrayL(key string, value interface{}) error
}
type Service struct {
	rabbitClient RabbitMQClient
	redisClient  RedisClient
}

func NewService(client RabbitMQClient, redisClient RedisClient) *Service {
	return &Service{rabbitClient: client, redisClient: redisClient}
}

func (s *Service) StartProcessing() {
	onReceivedChan := make(chan []byte, 0)
	go s.rabbitClient.Consume(onReceivedChan, Consumer)
	var forever chan struct{}
	go func() {
		for d := range onReceivedChan {
			var message Message
			err := json.Unmarshal(d, &message)
			if err != nil {
				continue
			}
			err = s.redisClient.SaveToArrayL(message.Sender, MessageDTO{
				Receiver: message.Receiver,
				Message:  message.Message,
			})
			if err != nil {
				fmt.Println("error occurred when save to array left ")
				continue
			}
			if err != nil {
				continue
			}
		}
	}()
	<-forever
}
