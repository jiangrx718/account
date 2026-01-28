package tags

import (
	"account/internal/service/tags"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	engine  *gin.Engine
	service tags.TagServiceIFace
}

func NewHandler(engine *gin.Engine) *Handler {
	return &Handler{
		engine:  engine,
		service: tags.New(),
	}
}

func (h *Handler) RegisterRoutes(routerGroup *gin.RouterGroup) {
	route := routerGroup.Group("/")
	route.Use()

	route.POST("/team/:team_id/tag_create", h.TagCreate) // 创建
}
