// Package errcode provides
package errcode

var (
	// ErrListBlocks .
	ErrListBlocks = NewErrcode(2000001, "列出區塊錯誤")
	// ErrGetBlock .
	ErrGetBlock = NewErrcode(2000002, "取得區塊錯誤")

	// ErrGetTransaction .
	ErrGetTransaction = NewErrcode(3000001, "取得交易錯誤")
)
