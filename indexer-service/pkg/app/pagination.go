// Package app provides
package app

import (
	"portto_interview/indexer-service/global"
	"portto_interview/indexer-service/pkg/convert"

	"github.com/gin-gonic/gin"
)

// Pagination .
type Pagination struct {
	Page      int `json:"page"`
	PageSize  int `json:"page_size"`
	TotalRows int `json:"total_rows"`
}

// Page .
func Page(c *gin.Context) int {
	page := convert.StrTo(c.Query("page")).MustInt()
	if page < 0 {
		return 1
	}
	return page
}

// PageSize .
func PageSize(c *gin.Context) int {
	size := convert.StrTo(c.Query("page_size")).MustInt()
	if size < 0 {
		return global.AppConfig.DefaultPageSize
	}
	if size > global.AppConfig.MaxPageSize {
		return global.AppConfig.MaxPageSize
	}
	return size
}

// PageOffset .
func PageOffset(page, size int) int {
	if page > 0 {
		return (page - 1) * size
	}
	return 0
}
