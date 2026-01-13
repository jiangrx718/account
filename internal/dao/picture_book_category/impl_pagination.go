package picture_book_category

import (
	"account/gopkg/gorms"
	"account/internal/dao"
	"account/internal/model"
	"context"
)

func (d *Dao) Pagination(ctx context.Context, page gorms.Page) (*gorms.Paging[*model.SPictureBookCategory], error) {
	paging, err := gorms.PaginationQuery(
		dao.SPictureBookCategory.Order(
			dao.SPictureBookCategory.Sort.Desc(),
		).FindByPage, gorms.Page{
			PageIndex: page.PageIndex,
			PageSize:  page.PageSize,
		})
	if err != nil {
		return nil, d.ConvertError(err)
	}

	return paging, nil
}
