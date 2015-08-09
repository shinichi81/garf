package echo

import (
	"github.com/backenderia/garf/server"
	"github.com/labstack/echo"
)

type echoGroup struct {
	*server.Default
	group *echo.Group
	echo  *echoHandler
}

// Group forwards to echo.Group
func (e *echoHandler) Group(d string) server.Router {
	return &echoGroup{
		group: e.echo.Group(d),
		echo:  e,
	}
}

// Get forwards to echo.Group.Use
func (g *echoGroup) Use(handler ...server.Middleware) {
	g.group.Use(handler)
}

// Get forwards to echo.Group.Get
func (g *echoGroup) Get(d string, handler server.HttpHandler) {
	g.group.Get(d, g.echo.contextWrapper(handler))
}

// Post forwards to echo.Group.Post
func (g *echoGroup) Post(d string, handler server.HttpHandler) {
	g.group.Post(d, g.echo.contextWrapper(handler))
}

// Put forwards to echo.Group.Put
func (g *echoGroup) Put(d string, handler server.HttpHandler) {
	g.group.Put(d, g.echo.contextWrapper(handler))
}

// Del forwards to echo.Group.Del
func (g *echoGroup) Del(d string, handler server.HttpHandler) {
	g.group.Delete(d, g.echo.contextWrapper(handler))
}

// Patch forwards to echo.Group.Patch
func (g *echoGroup) Patch(d string, handler server.HttpHandler) {
	g.group.Patch(d, g.echo.contextWrapper(handler))
}

// Options forwards to echo.Group.Options
func (g *echoGroup) Options(d string, handler server.HttpHandler) {
	g.group.Options(d, g.echo.contextWrapper(handler))
}

// Head forwards to echo.Group.Options
func (g *echoGroup) Head(d string, handler server.HttpHandler) {
	g.group.Head(d, g.echo.contextWrapper(handler))
}
