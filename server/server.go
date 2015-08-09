package server

import (
	"log"

	"golang.org/x/net/websocket"
)

// HttpHandler alias for any interface that handles requests
type HttpHandler func(Context) error
type Middleware interface{}

// Handler represents the HTTP server interface
type Handler interface {
	Configure()
	Run(string)
	Server() interface{}
	Group(string) Router
	Router
	ContextInfo
	RESTResponse
}

type Router interface {
	Use(...Middleware)
	Any(string, HttpHandler)
	Get(string, HttpHandler)
	Put(string, HttpHandler)
	Post(string, HttpHandler)
	Del(string, HttpHandler)
	Patch(string, HttpHandler)
	Options(string, HttpHandler)
	Head(string, HttpHandler)
	Handle(string, string, HttpHandler)
}

type RESTResponse interface {
	Error(Context, int, ...interface{}) error
	XML(Context, interface{}) error
	JSON(Context, interface{}) error
	HTML(Context, int, string, ...interface{}) error
	String(Context, int, string, ...interface{}) error
	NoContent(Context, int) error
}

// Default behavior for Server methods
type Default struct{}

// Configure default method
func (d *Default) Configure() { log.Println("Configure() not implemented on this server framework") }

// Server default method
func (d *Default) Server() interface{} {
	log.Println("Server() not implemented on this server framework")
	return nil
}

// Run default method
func (d *Default) Run(x string) { log.Println("Run() not implemented on this server framework") }

// Group default method
func (d *Default) Group(x string) { log.Println("Group() not implemented on this server framework") }

// Use default method
func (d *Default) Use(x Middleware) {
	log.Println("Use() not implemented on this server framework")
}

// Any default method
func (d *Default) Any(x string, y HttpHandler) {
	log.Println("Any() not implemented on this server framework")
}

// Get default method
func (d *Default) Get(x string, y HttpHandler) {
	log.Println("Get() not implemented on this server framework")
}

// Put default method
func (d *Default) Put(x string, y HttpHandler) {
	log.Println("Put() not implemented on this server framework")
}

// Post default method
func (d *Default) Post(x string, y HttpHandler) {
	log.Println("Post() not implemented on this server framework")
}

// Del default method
func (d *Default) Del(x string, y HttpHandler) {
	log.Println("Del() not implemented on this server framework")
}

// Patch default method
func (d *Default) Patch(x string, y HttpHandler) {
	log.Println("Patch() not implemented on this server framework")
}

// Options default method
func (d *Default) Options(x string, y HttpHandler) {
	log.Println("Options() not implemented on this server framework")
}

// Head default method
func (d *Default) Head(x string, y HttpHandler) {
	log.Println("Head() not implemented on this server framework")
}

// Handle default method
func (d *Default) Handle(x, y string, z HttpHandler) {
	log.Println("Handle() not implemented on this server framework")
}

// Param default method
func (d *Default) Param(x Context, y string) (z string) {
	log.Println("Param() not implemented on this server framework")
	return
}

// Param default method
func (d *Default) Form(x Context, y string) (z string) {
	log.Println("Form() not implemented on this server framework")
	return
}

// Query default method
func (d *Default) Query(x Context, y string) (z string) {
	log.Println("Query() not implemented on this server framework")
	return
}

// Error default method
func (d *Default) Error(x Context, y interface{}) (z error) {
	log.Println("Error() not implemented on this server framework")
	return
}

// JSON default method
func (d *Default) JSON(x Context, y interface{}) (z error) {
	log.Println("JSON() not implemented on this server framework")
	return
}

// String default method
func (d *Default) String(x Context, y int, z string, a ...interface{}) (b error) {
	log.Println("String() not implemented on this server framework")
	return
}

// HTML default method
func (d *Default) HTML(x Context, y int, z string, a ...interface{}) (b error) {
	log.Println("HTML() not implemented on this server framework")
	return
}

// XML default method
func (d *Default) XML(x Context, y interface{}) (z error) {
	log.Println("XML() not implemented on this server framework")
	return
}

// NoContent default method
func (d *Default) NoContent(x Context, y int) (z error) {
	log.Println("NoContent() not implemented on this server framework")
	return
}

// Socket default method
func (d *Default) Socket(x Context) (z *websocket.Conn) {
	log.Println("Socket() not implemented on this server framework")
	return
}
