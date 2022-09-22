// Package v1 provides
package v1

import (
	"portto_interview/api-service/internal/service"
	"portto_interview/api-service/pkg/app"
	"portto_interview/api-service/pkg/convert"
	"portto_interview/api-service/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// Block .
type Block struct{}

// NewBlock .
func NewBlock() *Block {
	return &Block{}
}

// List 列出區塊
func (b *Block) List(c *gin.Context) {
	params := service.GetBlockListReq{}
	err := c.ShouldBind(&params)
	if err != nil {
		app.NewResponse(c).ToError(errcode.InvalidParams)
		return
	}
	svc := service.NewService()
	blocks, err := svc.GetBlockList(c.Request.Context(), &params)
	if err != nil {
		app.NewResponse(c).ToError(errcode.ErrListBlocks)
		return
	}
	app.NewResponse(c).ToList(blocks, len(blocks))
	return
}

// Get 透過區塊ID取得單一區塊(包含所有 tx hash)
func (b *Block) Get(c *gin.Context) {
	params := service.GetBlockReq{
		ID: uint64(convert.StrTo(c.Param("id")).MustUInt64()),
	}
	err := c.ShouldBind(&params)
	if err != nil {
		app.NewResponse(c).ToError(errcode.InvalidParams)
		return
	}
	svc := service.NewService()
	block, err := svc.GetBlock(c.Request.Context(), &params)
	if err != nil {
		app.NewResponse(c).ToError(errcode.ErrGetBlock)
		return
	}
	app.NewResponse(c).ToResponse(block)
	return
}
