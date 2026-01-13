package sync

import (
	"account/gopkg/gins"
	"account/internal/service"
	"account/internal/service/sync"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	g           *gin.RouterGroup
	syncService service.Sync
}

func NewHandler(g *gin.RouterGroup) gins.Handler {
	return &Handler{
		g:           g,
		syncService: sync.NewService(),
	}
}

func (h *Handler) RegisterRoutes() {
	g := h.g.Group("/picture/sync")
	g.GET("/index", h.Index)
}
