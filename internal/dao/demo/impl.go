package demo

import (
	"account/gopkg/gorms"
)

type Dao struct {
	*gorms.BaseDao
}

func NewDao() *Dao {
	return &Dao{
		BaseDao: gorms.NewBaseDao(),
	}
}
