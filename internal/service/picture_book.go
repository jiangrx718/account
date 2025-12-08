package service

import (
	"account/gopkg/gorms"
	"account/gopkg/services"
	"context"
)

type PictureBook interface {
	PagingPictureBook(ctx context.Context, page gorms.Page) (services.Result, error)
}
