package tags

import (
	"account/internal/common"
	"context"
)

type TagServiceIFace interface {
	TagCreate(ctx context.Context, teamId, tagName, tagColor string, tagType int32) (common.ServiceResult, error)
}
