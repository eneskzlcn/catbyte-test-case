package message

import "encoding/json"

type RabbitMQClient interface {
	PushMessage(message []byte) error
}
type Service struct {
	client RabbitMQClient
}

func NewService(client RabbitMQClient) *Service {
	return &Service{client: client}
}

func (s *Service) PushMessage(message Message) error {
	messageBytes, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return s.client.PushMessage(messageBytes)
}
