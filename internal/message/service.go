package message

import (
	"encoding/json"
	"time"
)

type RabbitMQClient interface {
	PushMessage(message []byte) error
}
type Service struct {
	client RabbitMQClient
}

func NewService(client RabbitMQClient) *Service {
	return &Service{client: client}
}

func (s *Service) PushMessage(message *MessageRequest) error {
	rabbitMessage := message.ToRabbitMQMessage(time.Now())
	messageBytes, err := json.Marshal(rabbitMessage)
	if err != nil {
		return err
	}
	return s.client.PushMessage(messageBytes)
}
