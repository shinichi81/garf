package echo

import (
	"net/http"

	"github.com/backenderia/garf/server"
	"github.com/labstack/echo"
	"golang.org/x/net/websocket"
)

type echoContext struct {
	ctx    *echo.Context
	server server.Support
}

func NewEchoContext() *echoContext {
	return &echoContext{}
}

func (c *echoContext) Response() http.ResponseWriter {
	return c.ctx.Response()
}

func (c *echoContext) Request() *http.Request {
	return c.ctx.Request()
}

func (c *echoContext) Get(k string) interface{} {
	return c.ctx.Get(k)
}

func (c *echoContext) Set(k string, v interface{}) {
	c.ctx.Set(k, v)
}

func (c *echoContext) Context() interface{} {
	return c.ctx
}

func (c *echoContext) Server() server.Support {
	return c.server
}

// Param forwards to echo.Context.Param
func (c *echoContext) Param(d string) string {
	return c.ctx.Param(d)
}

// Form forwards to echo.Context.Form
func (c *echoContext) Form(d string) string {
	return c.ctx.Form(d)
}

// Query forwards to echo.Context.Query
func (c *echoContext) Query(d string) string {
	return c.ctx.Query(d)
}

// Socket forwards to echo.Context.Socket
func (c *echoContext) Socket() *websocket.Conn {
	return c.ctx.Socket()
}

// JSON forwards to echo.Context.JSON
func (c *echoContext) JSON(code int, d interface{}) error {
	return c.ctx.JSON(code, d)
}

// XML forwards to echo.Context.XML
func (c *echoContext) XML(code int, d interface{}) error {
	return c.ctx.XML(code, d)
}

// HTML forwards to echo.Context.HTML
func (c *echoContext) HTML(code int, data string, d ...interface{}) error {
	return c.ctx.HTML(code, data, d)
}

// String forwards to echo.Context.String
func (c *echoContext) String(code int, data string, d ...interface{}) error {
	return c.ctx.String(code, data, d)
}

// HTML forwards to echo.Context.HTML
func (c *echoContext) NoContent(d int) error {
	return c.ctx.NoContent(d)
}

// Error forwards to echo.Error
func (e *echoContext) Error(code int, d ...interface{}) error {
	if len(d) == 0 {
		return echo.NewHTTPError(code)
	}

	switch v := d[0].(type) {
	case error:
		return echo.NewHTTPError(code, v.Error())
	case string:
		return echo.NewHTTPError(code, v)
	default:
		return echo.NewHTTPError(code)
	}
}
