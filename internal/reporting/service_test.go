package reporting_test

import (
	mocks "github.com/eneskzlcn/catbyte-test-task/internal/mocks/reporting"
	"github.com/eneskzlcn/catbyte-test-task/internal/reporting"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewService(t *testing.T) {
	service := reporting.NewService(nil)
	assert.NotNil(t, service)
}
func TestService_Report(t *testing.T) {
	controller := gomock.NewController(t)
	mockRedisClient := mocks.NewMockRedisClient(controller)
	service := reporting.NewService(mockRedisClient)
	sender, receiver := "me", "you"
	mockRedisClient.EXPECT().GetArray(sender, gomock.Any()).Return(nil)
	_, err := service.Report(sender, receiver)
	assert.Nil(t, err)
}
