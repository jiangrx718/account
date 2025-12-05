package demo

import (
	"account/internal/dao"
	"account/internal/dao/demo"
)

type Service struct {
	demoDao dao.Demo
}

func NewService() *Service {
	return &Service{
		demoDao: demo.NewDao(),
	}
}
