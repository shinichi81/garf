package server

import (
	"log"

	"golang.org/x/net/websocket"
)

// Handler alias for any interface that handles requests
type Handler interface{}
type HandlerFunc interface{}
type Middleware interface{}

// Support represents the HTTP server interface
type Support interface {
	Configure()
	Run(string)
	Router
}

type Router interface {
	Server() interface{}
	Group(string) Router
	Use(...Middleware)
	Any(string, HandlerFunc)
	Get(string, HandlerFunc)
	Put(string, HandlerFunc)
	Post(string, HandlerFunc)
	Del(string, HandlerFunc)
	Patch(string, HandlerFunc)
	Options(string, HandlerFunc)
	Head(string, HandlerFunc)
	WebSocket(string, HandlerFunc)
	Handle(string, string, HandlerFunc)
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
func (d *Default) Any(x string, y HandlerFunc) {
	log.Println("Any() not implemented on this server framework")
}

// Get default method
func (d *Default) Get(x string, y HandlerFunc) {
	log.Println("Get() not implemented on this server framework")
}

// Put default method
func (d *Default) Put(x string, y HandlerFunc) {
	log.Println("Put() not implemented on this server framework")
}

// Post default method
func (d *Default) Post(x string, y HandlerFunc) {
	log.Println("Post() not implemented on this server framework")
}

// Del default method
func (d *Default) Del(x string, y HandlerFunc) {
	log.Println("Del() not implemented on this server framework")
}

// Patch default method
func (d *Default) Patch(x string, y HandlerFunc) {
	log.Println("Patch() not implemented on this server framework")
}

// Options default method
func (d *Default) Options(x string, y HandlerFunc) {
	log.Println("Options() not implemented on this server framework")
}

// Head default method
func (d *Default) Head(x string, y HandlerFunc) {
	log.Println("Head() not implemented on this server framework")
}

// WebSocket default method
func (d *Default) WebSocket(x, string, y HandlerFunc) {
	log.Println("WebSocket() not implemented on this server framework")
}

// Handle default method
func (d *Default) Handle(x, y string, z HandlerFunc) {
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
