package server

import (
	"net/http"

	"golang.org/x/net/websocket"
)

// Context interface
type Context interface {
	Request() *http.Request
	Response() http.ResponseWriter
	Set(string, interface{})
	Get(string) interface{}
}

// Context information
type ContextInfo interface {
	Param(Context, string) string
	Form(Context, string) string
	Query(Context, string) string
	Socket(Context) *websocket.Conn
}
