package service

import (
	"account/gopkg/gorms"
	"account/gopkg/services"
	"context"
)

type PictureBookCategory interface {
	PagingPictureBookCategory(ctx context.Context, page gorms.Page) (services.Result, error)
}
