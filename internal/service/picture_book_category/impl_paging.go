package picture_book_category

import (
	"account/gopkg/gorms"
	"account/gopkg/log"
	"account/gopkg/paging"
	"account/gopkg/services"
	"account/internal/model"
	"context"

	"go.uber.org/zap"
)

func (s *Service) PagingPictureBookCategory(ctx context.Context, page gorms.Page) (services.Result, error) {
	logPrefix := "/internal/service/picture_book_category: Service.PagingPictureBookCategory()"

	demoPaging, err := s.pictureBookCategoryDao.Pagination(ctx, page)
	if err != nil {
		log.Sugar().Error(logPrefix, zap.Any("picture_book_category dao pagination error", err), zap.Any("page", page))
		return services.Failed(ctx, err)
	}
	return services.Success(ctx, paging.NewPaging(demoPaging.Total, NewPictureBookCategoryS(demoPaging.List)))
}

func NewPictureBookCategoryS(demoEntities []*model.SPictureBookCategory) []*model.SPictureBookCategory {

	return demoEntities
}
