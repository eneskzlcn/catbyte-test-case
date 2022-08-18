package message

import "time"

type MessageRequest struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Message  string `json:"message"`
}

type RabbitMQMessage struct {
	Sender   string    `json:"sender"`
	Receiver string    `json:"receiver"`
	Message  string    `json:"message"`
	Date     time.Time `json:"date"`
}

func (m *MessageRequest) ToRabbitMQMessage(time time.Time) RabbitMQMessage {
	return RabbitMQMessage{
		Sender:   m.Sender,
		Receiver: m.Receiver,
		Message:  m.Message,
		Date:     time,
	}
}
