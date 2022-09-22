// Package service provides
package service

import (
	"portto_interview/api-service/global"
	"portto_interview/api-service/internal/dao"
)

// Service .
type Service struct {
	dao *dao.Dao
}

// NewService .
func NewService() *Service {
	return &Service{
		dao: dao.NewDao(global.Database),
	}
}
