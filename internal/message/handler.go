package message

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type MessageService interface {
	PushMessage(message *MessageRequest) error
}

func NewHandler(service MessageService) *Handler {
	return &Handler{service: service}
}

type Handler struct {
	service MessageService
}

func (h *Handler) GetMessage(ctx *gin.Context) {
	var message MessageRequest
	if err := ctx.BindJSON(&message); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}
	if err := h.service.PushMessage(&message); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}
	ctx.Status(http.StatusOK)
}

func (h *Handler) RegisterRoutes(engine *gin.Engine) {
	engine.GET("/message", h.GetMessage)
}
