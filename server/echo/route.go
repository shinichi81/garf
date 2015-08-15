package echo

import (
	"strings"

	"github.com/backenderia/garf/server"
	"github.com/labstack/echo"
)

func (e *echoHandler) contextWrap(h server.HandlerFunc) echo.HandlerFunc {
	switch h := h.(type) {
	case func(server.Context) error:
		return func(c *echo.Context) (result error) {
			ec := e.pool.Get().(*echoContext)
			ec.ctx = c
			ec.server = e
			result = h(ec)
			e.pool.Put(ec)
			return
		}
	default:
		return h.(echo.HandlerFunc)
	}
}

func (e *echoHandler) middlewareWrap(m server.Middleware) server.Middleware {
	switch m := m.(type) {
	case func(server.Context) error:
		return e.contextWrap(m)
	default:
		return m
	}
}

// Get forwards to echo.Use
func (e *echoHandler) Use(m ...server.Middleware) {
	em := []echo.Middleware{}
	for _, h := range m {
		em = append(em, e.middlewareWrap(h))
	}
	e.echo.Use(em...)
}

// Get forwards to echo.Get
func (e *echoHandler) Get(d string, handler server.HandlerFunc) {
	e.echo.Get(d, e.contextWrap(handler))
}

// Post forwards to echo.Post
func (e *echoHandler) Post(d string, handler server.HandlerFunc) {
	e.echo.Post(d, e.contextWrap(handler))
}

// Put forwards to echo.Put
func (e *echoHandler) Put(d string, handler server.HandlerFunc) {
	e.echo.Put(d, e.contextWrap(handler))
}

// Del forwards to echo.Del
func (e *echoHandler) Del(d string, handler server.HandlerFunc) {
	e.echo.Delete(d, e.contextWrap(handler))
}

// Patch forwards to echo.Patch
func (e *echoHandler) Patch(d string, handler server.HandlerFunc) {
	e.echo.Patch(d, e.contextWrap(handler))
}

// Options forwards to echo.Options
func (e *echoHandler) Options(d string, handler server.HandlerFunc) {
	e.echo.Options(d, e.contextWrap(handler))
}

// Head forwards to echo.Head
func (e *echoHandler) Head(d string, handler server.HandlerFunc) {
	e.echo.Head(d, e.contextWrap(handler))
}

// WebSocket forwards to echo.WebSocket
func (e *echoHandler) WebSocket(d string, handler server.HandlerFunc) {
	e.echo.WebSocket(d, e.contextWrap(handler))
}

// Handle forwards to echo.{Method}
func (e *echoHandler) Handle(c, d string, handler server.HandlerFunc) {
	switch strings.ToUpper(c) {
	case "GET":
		e.Get(d, handler)
	case "PUT":
		e.Put(d, handler)
	case "POST":
		e.Post(d, handler)
	case "DEL":
		e.Del(d, handler)
	case "DELETE":
		e.Del(d, handler)
	case "PATCH":
		e.Patch(d, handler)
	case "OPTIONS":
		e.Options(d, handler)
	case "HEAD":
		e.Head(d, handler)
	case "WEBSOCKET":
		e.WebSocket(d, handler)
	}
}
