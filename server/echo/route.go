package echo

import "github.com/backenderia/garf/server"

// Get forwards to echo.Use
func (e *echoHandler) Use(handler server.HttpHandler) {
	e.echo.Use(e.contextWrapper(handler))
}

// Get forwards to echo.Get
func (e *echoHandler) Get(d string, handler server.HttpHandler) {
	e.echo.Get(d, e.contextWrapper(handler))
}

// Post forwards to echo.Post
func (e *echoHandler) Post(d string, handler server.HttpHandler) {
	e.echo.Post(d, e.contextWrapper(handler))
}

// Put forwards to echo.Put
func (e *echoHandler) Put(d string, handler server.HttpHandler) {
	e.echo.Put(d, e.contextWrapper(handler))
}

// Del forwards to echo.Del
func (e *echoHandler) Del(d string, handler server.HttpHandler) {
	e.echo.Delete(d, e.contextWrapper(handler))
}

// Patch forwards to echo.Patch
func (e *echoHandler) Patch(d string, handler server.HttpHandler) {
	e.echo.Patch(d, e.contextWrapper(handler))
}

// Options forwards to echo.Options
func (e *echoHandler) Options(d string, handler server.HttpHandler) {
	e.echo.Options(d, e.contextWrapper(handler))
}

// Head forwards to echo.Options
func (e *echoHandler) Head(d string, handler server.HttpHandler) {
	e.echo.Head(d, e.contextWrapper(handler))
}
