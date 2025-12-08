package dao

import (
	"account/gopkg/gorms"
	"account/internal/model"
	"context"
)

type PictureBookCategory interface {
	Pagination(ctx context.Context, page gorms.Page) (*gorms.Paging[*model.SPictureBookCategory], error)
}
