package echo_test

import (
	"fmt"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/backenderia/garf/server"
	_echo "github.com/backenderia/garf/server/echo"
	"github.com/labstack/echo"
	. "github.com/smartystreets/goconvey/convey"
	"golang.org/x/net/websocket"
)

func TestEchoSupport(t *testing.T) {
	Convey("Given a new garf's echo server...", t, func() {
		s := _echo.New()
		Convey("The method echo.Server() should return a valid *echo.Echo", func() {
			_, isEcho := (s.Server()).(*echo.Echo)
			So(isEcho, ShouldBeTrue)
		})
	})
}

func TestEchoSupportParam(t *testing.T) {
	Convey("Given a echo.Server", t, func() {
		s := _echo.New()
		e := (s.Server()).(*echo.Echo)
		Convey("When adding a handler to a request with parameters", func() {
			param := ""
			s.Get("/:param", func(c server.Context) error {
				param = s.Param(c, "param")
				return nil
			})
			request("GET", "/123", e)
			Convey("The response Param() should be 123", func() {
				So(param, ShouldEqual, "123")
			})
		})
	})
}

func TestEchoSupportForm(t *testing.T) {
	Convey("Given a echo.Server", t, func() {
		s := _echo.New()
		e := (s.Server()).(*echo.Echo)
		Convey("When adding a handler to a request using form", func() {
			f := make(url.Values)
			f.Set("form", "123")

			form := ""
			s.Post("/", func(c server.Context) error {
				form = s.Form(c, "form")
				return nil
			})

			requestPost("POST", "/", strings.NewReader(f.Encode()), e)

			Convey("The response Form() should be 123", func() {
				So(form, ShouldEqual, "123")
			})
		})
	})
}

func TestEchoSupportQuery(t *testing.T) {
	Convey("Given a echo.Server", t, func() {
		s := _echo.New()
		e := (s.Server()).(*echo.Echo)
		Convey("When adding a handler to a request with query", func() {
			query := ""
			s.Get("/", func(c server.Context) error {
				query = s.Query(c, "q")
				return nil
			})
			request("GET", "/?q=123", e)
			Convey("The response Query() should be 123", func() {
				So(query, ShouldEqual, "123")
			})
		})
	})
}

func TestEchoSupportSocket(t *testing.T) {
	Convey("Given a echo.Server", t, func() {
		s := _echo.New()
		e := (s.Server()).(*echo.Echo)
		Convey("An attached handler to a WebSocket request that calls server.Socket(context)", func() {
			var sock *websocket.Conn
			s.WebSocket("/ws", func(c server.Context) error {
				sock = s.Socket(c)
				return nil
			})
			srv := httptest.NewServer(e)
			defer srv.Close()
			addr := srv.Listener.Addr().String()
			origin := "http://localhost"
			url := fmt.Sprintf("ws://%s/ws", addr)
			websocket.Dial(url, "", origin)
			Convey("Socket() should return a valid *websocket.Conn", func() {
				So(sock, ShouldNotBeNil)
			})
		})
	})
}
