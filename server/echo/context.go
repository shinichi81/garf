package echo

import (
	"net/http"

	"github.com/labstack/echo"
)

type echoContext struct {
	ctx *echo.Context
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
