package reporting_test

import (
	"encoding/json"
	"fmt"
	mocks "github.com/eneskzlcn/catbyte-test-task/internal/mocks/reporting"
	"github.com/eneskzlcn/catbyte-test-task/internal/reporting"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

func TestNewHandler(t *testing.T) {
	handler := reporting.NewHandler(nil)
	assert.NotNil(t, handler)
}
func TestHandler_Report(t *testing.T) {
	controller := gomock.NewController(t)
	mockReportingService := mocks.NewMockReportingService(controller)
	handler := reporting.NewHandler(mockReportingService)
	assert.NotNil(t, handler)
	engine := gin.Default()
	engine.SetTrustedProxies([]string{"*"})
	engine.GET("/message/list", handler.Report)

	url, route, sender, receiver := "http://localhost:8080", "/message/list", "me", "you"
	expectedReports := []reporting.Report{
		{
			Sender:   sender,
			Receiver: receiver,
			Message:  "hi",
		},
		{
			Sender:   sender,
			Receiver: receiver,
			Message:  "you",
		},
		{
			Sender:   sender,
			Receiver: receiver,
			Message:  "3",
		},
	}
	mockReportingService.EXPECT().Report(sender, receiver).Return(expectedReports, nil)
	go engine.Run("localhost:8080")
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s?sender=%s&receiver=%s", url, route, sender, receiver), nil)
	assert.Nil(t, err)
	resp, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	actualReports := make([]reporting.Report, 0)
	err = json.Unmarshal(body, &actualReports)
	assert.Nil(t, err)
	assert.ElementsMatch(t, expectedReports, actualReports)
}
