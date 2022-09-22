// Package app provides
package app

import (
	"net/http"
	"portto_interview/indexer-service/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// Response .
type Response struct {
	_ctx *gin.Context
}

// ctx .
func (r *Response) ctx() *gin.Context {
	return r._ctx
}

// NewResponse .
func NewResponse(ctx *gin.Context) *Response {
	return &Response{ctx}
}

// ToResponse .
func (r *Response) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	r.ctx().JSON(http.StatusOK, data)
}

// ToList .
func (r *Response) ToList(data interface{}, rows int) {
	if data == nil {
		data = gin.H{}
	}
	if rows <= 0 {
		rows = 0 // 可以處理分頁需求
	}
	ctx := r.ctx()
	ctx.JSON(http.StatusOK, gin.H{
		"data": data,
		// Add pagination at here
		"pager": Pagination{
			Page:      Page(ctx),
			PageSize:  PageSize(ctx),
			TotalRows: rows,
		},
	})
}

// ToError .
func (r *Response) ToError(err *errcode.Error) {
	r.ctx().JSON(err.StatusCode(), gin.H{
		"code": err.Code(),
		"msg":  err.Msg(),
	})
}
