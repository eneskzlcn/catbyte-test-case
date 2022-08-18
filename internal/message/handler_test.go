package message_test

import (
	"bytes"
	"encoding/json"
	message "github.com/eneskzlcn/catbyte-test-task/internal/message"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewHandler(t *testing.T) {
	handler := message.NewHandler(nil)
	assert.NotNil(t, handler)
}
func TestHandler_GetMessage(t *testing.T) {
	t.Run("given valid request then it should return success", func(t *testing.T) {
		handler := message.NewHandler(nil)
		router := gin.Default()
		router.SetTrustedProxies([]string{"*"})

		router.GET("/message", handler.GetMessage)
		go router.Run("localhost:8080")
		message := message.MessageRequest{
			Sender:   "me",
			Receiver: "you",
			Message:  "hey",
		}
		messageBytes, err := json.Marshal(message)
		assert.Nil(t, err)
		req, err := http.NewRequest("GET", "http://localhost:8080/message", bytes.NewReader(messageBytes))
		assert.Nil(t, err)
		resp, err := http.DefaultClient.Do(req)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

	})
	t.Run("given invalid request then it should return bad request", func(t *testing.T) {
		handler := message.NewHandler(nil)
		router := gin.Default()
		router.SetTrustedProxies([]string{"*"})
		go router.Run("localhost:8080")
		router.GET("/message", handler.GetMessage)
		req, err := http.NewRequest("GET", "http://localhost:8080/message", nil)
		assert.Nil(t, err)
		resp, err := http.DefaultClient.Do(req)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})
}
