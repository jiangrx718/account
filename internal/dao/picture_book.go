package dao

import (
	"account/gopkg/gorms"
	"account/internal/model"
	"context"
)

type PictureBook interface {
	Pagination(ctx context.Context, page gorms.Page) (*gorms.Paging[*model.SPictureBook], error)
}
