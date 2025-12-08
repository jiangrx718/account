package picture_book

import (
	"account/gopkg/gins"
	"account/internal/service"
	"account/internal/service/picture_book"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	g                  *gin.RouterGroup
	pictureBookService service.PictureBook
}

func NewHandler(g *gin.RouterGroup) gins.Handler {
	return &Handler{
		g:                  g,
		pictureBookService: picture_book.NewService(),
	}
}

func (h *Handler) RegisterRoutes() {
	g := h.g.Group("/picture/book")
	g.GET("/list", h.PagingPictureBook)
}
