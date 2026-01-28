package service

import (
	"account/internal/common"
)

func NewServiceResult() *ServiceResult {
	return &ServiceResult{}
}

type ServiceResult struct {
	common.BaseServiceResult
	Offset int64 `json:"offset,omitempty"`
	Limit  int64 `json:"limit,omitempty"`
	Count  int64 `json:"count,omitempty"`
}
