package api

import (
	"account/gopkg/gins"
	"account/handler/api/demo"
	"account/handler/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	engine *gin.Engine
}

func NewHandler(engine *gin.Engine) gins.Handler {
	return &Handler{
		engine: engine,
	}
}

func (h *Handler) RegisterRoutes() {
	config := cors.DefaultConfig()
	config.AllowHeaders = append([]string{}, config.AllowHeaders...)
	config.AllowAllOrigins = true
	h.engine.Use(cors.New(config))

	g := h.engine.Group("/api", middleware.RequestCapture())
	handlers := []gins.Handler{
		demo.NewHandler(g),
	}

	for _, handler := range handlers {
		handler.RegisterRoutes()
	}
}
