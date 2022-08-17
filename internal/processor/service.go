package processor

import "encoding/json"

type RabbitMQClient interface {
	Consume(onRecieved chan<- []byte)
}
type RedisClient interface {
	SaveStruct(key string, value interface{}) error
}
type Service struct {
	rabbitClient RabbitMQClient
	redisClient  RedisClient
}

func NewService(client RabbitMQClient, redisClient RedisClient) *Service {
	return &Service{rabbitClient: client, redisClient: redisClient}
}

func (s *Service) StartProcessing() {
	onRecievedChan := make(chan<- []byte, 0)
	s.rabbitClient.Consume(onRecievedChan)
	var forever chan struct{}
	go func() {
		for d := range onRecievedChan {
			var message Message
			err := json.Unmarshal(d, &message)
			if err != nil {
				continue
			}
			err = s.redisClient.SaveStruct(message.Sender, message)
			if err != nil {
				continue
			}
		}
	}()
	<-forever
}
