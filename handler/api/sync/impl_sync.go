package sync

import (
	"account/gopkg/gins"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Index(ctx *gin.Context) {

	res, err := h.syncService.Index(ctx)
	if err != nil {
		gins.ServerError(ctx, err)
		return
	}

	gins.StatusOK(ctx, res)
}
