// Package v1 provides
package v1

import (
	"portto_interview/api-service/internal/service"
	"portto_interview/api-service/pkg/app"
	"portto_interview/api-service/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// Transaction .
type Transaction struct{}

// NewTransaction .
func NewTransaction() *Transaction {
	return &Transaction{}
}

// Get 透過 Tx Hash 取得單一交易(包含event log)
func (t *Transaction) Get(c *gin.Context) {
	params := service.GetTransactionReq{
		TxHash: c.Param("txHash"),
	}
	err := c.ShouldBind(&params)
	if err != nil {
		app.NewResponse(c).ToError(errcode.InvalidParams)
		return
	}
	svc := service.NewService()
	transaction, err := svc.GetTransaction(c.Request.Context(), &params)
	if err != nil {
		app.NewResponse(c).ToError(errcode.ErrGetTransaction)
		return
	}
	app.NewResponse(c).ToResponse(transaction)
	return
}
