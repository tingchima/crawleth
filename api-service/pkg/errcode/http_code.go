// Package errcode provides
package errcode

import "net/http"

// StatusCode .
func (e *Error) StatusCode() int {
	switch e.code {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}
