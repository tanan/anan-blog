package handler

import "net/http"

type Context interface {
	Param(string) string
	QueryParam(name string) string
	String(int, string) error
	JSON(code int, i interface{}) error
	Bind(i interface{}) error
	Request() *http.Request
}
