package picture_book_category

import (
	"account/gopkg/gorms"
	"account/internal/g"
	"account/internal/model"
	"context"
)

func (d *Dao) Pagination(ctx context.Context, page gorms.Page) (*gorms.Paging[*model.SPictureBookCategory], error) {
	paging, err := gorms.PaginationQuery(
		g.SPictureBookCategory.Order(
			g.SPictureBookCategory.CreatedAt.Desc(),
		).FindByPage, gorms.Page{
			PageIndex: page.PageIndex,
			PageSize:  page.PageSize,
		})
	if err != nil {
		return nil, d.ConvertError(err)
	}

	return paging, nil
}
