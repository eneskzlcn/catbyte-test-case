package message

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type MessageService interface {
	PushMessage(message *Message) error
}

func NewHandler(service MessageService) *Handler {
	return &Handler{service: service}
}

type Handler struct {
	service MessageService
}

func (h *Handler) GetMessage(ctx *gin.Context) {
	var message Message
	if err := ctx.BindJSON(&message); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}
	if err := h.service.PushMessage(&message); err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.Status(http.StatusOK)
}

func (h *Handler) RegisterRoutes(engine *gin.Engine) {
	engine.GET("/message", h.GetMessage)
}
