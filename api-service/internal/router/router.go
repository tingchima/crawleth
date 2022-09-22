// Package router provides
package router

import (
	v1 "portto_interview/api-service/internal/router/api/v1"

	"github.com/gin-gonic/gin"
)

// NewRouter .
func NewRouter() *gin.Engine {
	r := gin.New()

	// Middleware can add some custom middleware at here

	// Routers
	block := v1.NewBlock()
	transaction := v1.NewTransaction()

	// Blocks
	r.GET("/blocks", block.List)
	r.GET("/blocks/:id", block.Get)
	// Transactions
	r.GET("/transactions/:txHash", transaction.Get)

	return r
}
