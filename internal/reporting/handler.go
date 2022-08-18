package reporting

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ReportingService interface {
	Report(sender, receiver string) ([]Report, error)
}
type Handler struct {
	service ReportingService
}

func NewHandler(service ReportingService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Report(ctx *gin.Context) {
	sender := ctx.Query("sender")
	receiver := ctx.Query("receiver")
	report, err := h.service.Report(sender, receiver)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, report)
}
func (h *Handler) RegisterRoutes(engine *gin.Engine) {
	engine.GET("/message/list", h.Report)
}
