package picture_book

import (
	"account/internal/dao"
	"account/internal/dao/picture_book"
)

type Service struct {
	pictureBookDao dao.PictureBook
}

func NewService() *Service {
	return &Service{
		pictureBookDao: picture_book.NewDao(),
	}
}
