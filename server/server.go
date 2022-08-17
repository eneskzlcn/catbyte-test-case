package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler interface {
	RegisterRoutes(engine *gin.Engine)
}
type Server struct {
	engine *gin.Engine
	port   string
}

func NewServer(handlers []Handler, trustedProxies []string, port string) *Server {
	engine := gin.Default()
	engine.SetTrustedProxies(trustedProxies)
	server := Server{port: port, engine: engine}
	for _, handler := range handlers {
		handler.RegisterRoutes(server.engine)
	}
	server.AddRoutes()
	return &server
}
func (s *Server) AddRoutes() {
	s.engine.GET("/health", s.healthCheck)
}
func (s *Server) healthCheck(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}
func (s *Server) Start() error {
	return s.engine.Run(fmt.Sprintf("localhost:%s", s.port))
}
