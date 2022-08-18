package message_test

import (
	"encoding/json"
	"github.com/eneskzlcn/catbyte-test-task/internal/message"
	mocks "github.com/eneskzlcn/catbyte-test-task/internal/mocks/message"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewService(t *testing.T) {
	service := message.NewService(nil)
	assert.NotNil(t, service)
}
func TestService_PushMessage(t *testing.T) {
	controller := gomock.NewController(t)
	mockRabbitMQClient := mocks.NewMockRabbitMQClient(controller)
	service := message.NewService(mockRabbitMQClient)

	testMessage := message.MessageRequest{
		Sender:   "me",
		Receiver: "you",
		Message:  "hey",
	}
	rabbitMessage := message.RabbitMQMessage{
		Sender:   testMessage.Sender,
		Receiver: testMessage.Receiver,
		Message:  testMessage.Message,
	}

	messageBytes, err := json.Marshal(testMessage)
	assert.Nil(t, err)
	mockRabbitMQClient.EXPECT().PushMessage(messageBytes).Return(nil)
	err = service.PushMessage(&testMessage, rabbitMessage.Date)
	assert.Equal(t, err, nil)
}
