package service

import (
	"account/gopkg/services"
	"context"
)

type Sync interface {
	Index(ctx context.Context) (services.Result, error)
}
