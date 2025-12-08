package picture_book_category

import (
	"account/internal/dao"
	"account/internal/dao/picture_book_category"
)

type Service struct {
	pictureBookCategoryDao dao.PictureBookCategory
}

func NewService() *Service {
	return &Service{
		pictureBookCategoryDao: picture_book_category.NewDao(),
	}
}
