package echo

import (
	"sync"

	"github.com/backenderia/garf/server"
	"github.com/labstack/echo"
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
