package echo

import (
	"sync"

	"github.com/backenderia/garf/server"
	"github.com/labstack/echo"
	"golang.org/x/net/websocket"
)

type echoHandler struct {
	*server.Default
	pool   sync.Pool
	echo   *echo.Echo
	prefix string
}

// New creates a new Registry for this server
func New() server.Support {
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

// Group creates a new echoGroup
func (e *echoHandler) Group(d string) server.Router {
	group := &echoGroup{
		group: e.echo.Group(d),
	}
	group.pool.New = func() interface{} {
		return NewEchoContext()
	}
	return group
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
