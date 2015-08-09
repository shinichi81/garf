package echo

import (
	"sync"

	"github.com/backenderia/garf/server"
	"github.com/labstack/echo"
	"golang.org/x/net/websocket"
)

type echoHandler struct {
	*server.Default
	pool sync.Pool
	echo *echo.Echo
}

// New creates a new Registry for this server
func New() server.Handler {
	e := echo.New()

	echo := &echoHandler{
		echo: e,
	}

	echo.pool.New = func() interface{} {
		return NewEchoContext()
	}
	return echo
}

// Run forwards to echo.Run
func (e *echoHandler) Run(d string) {
	e.echo.Run(d)
}

// Configure the echo server
func (e *echoHandler) Configure() {
}

// Server returns the echo instance
func (e *echoHandler) Server() interface{} {
	return e.echo
}

func (e *echoHandler) contextWrapper(h server.HttpHandler) func(*echo.Context) error {
	return func(c *echo.Context) (result error) {
		ec := e.pool.Get().(*echoContext)
		ec.ctx = c
		result = h(ec)
		e.pool.Put(ec)
		return
	}
}

// Param forwards to echo.Context.Param
func (e *echoHandler) Param(c server.Context, d string) string {
	ec := c.(*echoContext)
	return ec.ctx.Param(d)
}

// Form forwards to echo.Context.Form
func (e *echoHandler) Form(c server.Context, d string) string {
	ec := c.(*echoContext)
	return ec.ctx.Form(d)
}

// Query forwards to echo.Context.Query
func (e *echoHandler) Query(c server.Context, d string) string {
	ec := c.(*echoContext)
	return ec.ctx.Query(d)
}

// File forwards to echo.Context.File
func (e *echoHandler) Socket(c server.Context) *websocket.Conn {
	ec := c.(*echoContext)
	return ec.ctx.Socket()
}
