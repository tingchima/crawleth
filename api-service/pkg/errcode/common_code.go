// Package errcode provides
package errcode

var (
	// Success .
	Success = NewErrcode(0, "成功")
	// ServerError .
	ServerError = NewErrcode(1000000, "服務內部錯誤")
	// InvalidParams .
	InvalidParams = NewErrcode(1000001, "參數錯誤")
	// NotFound .
	NotFound = NewErrcode(1000002, "找不到資源")
	// Add new custom error code at here ...
)
