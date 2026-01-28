package tags

import (
	"account/server/http/httputil"

	"github.com/gin-gonic/gin"
)

type TagUri struct {
	TeamID string `uri:"team_id" binding:"required"`
}

type TagCreateReq struct {
	TagName  string `json:"tag_name" binding:"required"`
	TagColor string `json:"tag_color" binding:"required"`
	TagType  int    `json:"tag_type" binding:"required"`
}

// TagCreate			创建Tag
// @Tags				标签接口
// @Summary				创建Tag
// @Description			创建Tag
// @Accept	json
// @Produce	json
// @Param Authorization header string true "Bearer {token}"
// @Param path TagUri true "uri参数"
// @Success 400 {object} common.BaseServiceResult
// @Success 200 {object} studio.StudioServiceResult{}
// @Router /team/:team_id/tag_create [post]
func (h *Handler) TagCreate(ctx *gin.Context) {
	var req TagUri
	if err := ctx.ShouldBindUri(&req); err != nil {
		httputil.BadRequest(ctx, err)
		return
	}

	var reqBody TagCreateReq
	if err := ctx.Bind(&reqBody); err != nil {
		httputil.BadRequest(ctx, err)
		return
	}

	result, err := h.service.TagCreate(ctx, req.TeamID, reqBody.TagName, reqBody.TagColor, int32(reqBody.TagType))
	if err != nil {
		httputil.ServerError(ctx, err)
		return
	}

	ctx.JSON(200, result)
	return
}
