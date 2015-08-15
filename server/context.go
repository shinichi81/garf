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
	Context() interface{}
	Server() Support
	ContextInfo
	RESTResponse
}

// ContextInfo methods
type ContextInfo interface {
	Param(string) string
	Form(string) string
	Query(string) string
	Socket() *websocket.Conn
}

// RESTResponse methods
type RESTResponse interface {
	Error(int, ...interface{}) error
	XML(int, interface{}) error
	JSON(int, interface{}) error
	HTML(int, string, ...interface{}) error
	String(int, string, ...interface{}) error
	NoContent(int) error
}
