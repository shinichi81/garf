package echo

import (
	"strings"
	"sync"

	"github.com/backenderia/garf/server"
	"github.com/labstack/echo"
)

type echoGroup struct {
	*server.Default
	pool  sync.Pool
	group *echo.Group
}

func (g *echoGroup) contextWrap(h server.HandlerFunc) echo.HandlerFunc {
	switch h := h.(type) {
	case func(server.Context) error:
		return func(c *echo.Context) (result error) {
			ec := g.pool.Get().(*echoContext)
			ec.ctx = c
			ec.server = g
			result = h(ec)
			g.pool.Put(ec)
			return
		}
	default:
		return h.(echo.HandlerFunc)
	}
}

func (g *echoGroup) middlewareWrap(m server.Middleware) server.Middleware {
	switch m := m.(type) {
	case func(server.Context) error:
		return g.contextWrap(m)
	default:
		return m
	}
}

// Group forwards to echo.Group
func (g *echoGroup) Group(d string) server.Router {
	group := &echoGroup{
		group: g.group.Group(d),
	}
	group.pool.New = func() interface{} {
		return NewEchoContext()
	}
	return group
}

// Server returns a *echo.Group
func (g *echoGroup) Server() interface{} {
	return g.group
}

// Get forwards to echo.Group.Use
func (g *echoGroup) Use(m ...server.Middleware) {
	em := []echo.Middleware{}
	for _, h := range m {
		em = append(em, g.middlewareWrap(h))
	}
	g.group.Use(em...)
}

// Get forwards to echo.Group.Get
func (g *echoGroup) Get(d string, handler server.HandlerFunc) {
	g.group.Get(d, g.contextWrap(handler))
}

// Post forwards to echo.Group.Post
func (g *echoGroup) Post(d string, handler server.HandlerFunc) {
	g.group.Post(d, g.contextWrap(handler))
}

// Put forwards to echo.Group.Put
func (g *echoGroup) Put(d string, handler server.HandlerFunc) {
	g.group.Put(d, g.contextWrap(handler))
}

// Del forwards to echo.Group.Del
func (g *echoGroup) Del(d string, handler server.HandlerFunc) {
	g.group.Delete(d, g.contextWrap(handler))
}

// Patch forwards to echo.Group.Patch
func (g *echoGroup) Patch(d string, handler server.HandlerFunc) {
	g.group.Patch(d, g.contextWrap(handler))
}

// Options forwards to echo.Group.Options
func (g *echoGroup) Options(d string, handler server.HandlerFunc) {
	g.group.Options(d, g.contextWrap(handler))
}

// Head forwards to echo.Group.Head
func (g *echoGroup) Head(d string, handler server.HandlerFunc) {
	g.group.Head(d, g.contextWrap(handler))
}

// WebSocket forwards to echo.Group.WebSocket
func (g *echoGroup) WebSocket(d string, handler server.HandlerFunc) {
	g.group.WebSocket(d, g.contextWrap(handler))
}

// Handle forwards to echo.Group.{Method}
func (g *echoGroup) Handle(c, d string, handler server.HandlerFunc) {
	switch strings.ToUpper(c) {
	case "GET":
		g.Get(d, handler)
	case "PUT":
		g.Put(d, handler)
	case "POST":
		g.Post(d, handler)
	case "DEL":
		g.Del(d, handler)
	case "DELETE":
		g.Del(d, handler)
	case "PATCH":
		g.Patch(d, handler)
	case "OPTIONS":
		g.Options(d, handler)
	case "HEAD":
		g.Head(d, handler)
	case "WEBSOCKET":
		g.WebSocket(d, handler)
	}
}
