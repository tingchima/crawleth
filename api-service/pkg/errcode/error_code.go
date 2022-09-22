// Package errcode provides
package errcode

import "fmt"

var errcodes = map[int]*Error{}

// Error .
type Error struct {
	code int    // 錯誤代號
	msg  string // 錯誤訊息
}

// NewErrcode .
func NewErrcode(code int, msg string) *Error {
	_, exist := errcodes[code]
	if exist {
		panic(fmt.Sprintf("%d error code 已經存在", code))
	}
	var err = Error{
		code: code,
		msg:  msg,
	}
	errcodes[code] = &err
	return &err
}

// Error .
func (e *Error) Error() string {
	return fmt.Sprintf("錯誤: %d, 錯誤訊息:%s", e.code, e.msg)
}

// Code .
func (e *Error) Code() int {
	return e.code
}

// Msg .
func (e *Error) Msg() string {
	return e.msg
}

// Msgf .
func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.msg, args...)
}
