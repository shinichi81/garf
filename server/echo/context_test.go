package echo_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
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

func TestEchoSupportContext(t *testing.T) {
	Convey("Given a echo.Server", t, func() {
		s := _echo.New()
		e := (s.Server()).(*echo.Echo)
		Convey(`A handler should be able to context.Set("n", 123)`, func() {
			var n, a int
			s.Use(func(c server.Context) error {
				ctx := c.Context().(*echo.Context)
				ctx.Set("a", 123)
				c.Set("n", 123)
				return nil
			})
			s.Use(func(c server.Context) error {
				n = (c.Get("n")).(int)
				ctx := c.Context().(*echo.Context)
				a = ctx.Get("a").(int)
				return nil
			})
			request("GET", "/", e)
			Convey(`context.Get("n") should return 123`, func() {
				So(n, ShouldEqual, 123)
				So(a, ShouldEqual, 123)
			})
		})
	})
}

func TestEchoSupportGroupContext(t *testing.T) {
	Convey("Given a echo.Server", t, func() {
		srv := _echo.New()
		e := (srv.Server()).(*echo.Echo)
		s := srv.Group("/123")
		Convey(`A group route handler should be able to context.Set("n", 123)`, func() {
			var n int
			s.Use(func(c server.Context) error {
				c.Set("n", 123)
				return nil
			})
			s.Use(func(c server.Context) error {
				n = (c.Get("n")).(int)
				return nil
			})
			s.Get("/", func(c server.Context) error {
				return nil
			})
			request("GET", "/123/", e)
			Convey(`context.Get("n") should return 123`, func() {
				So(n, ShouldEqual, 123)
			})
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
				param = c.Param("param")
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
				form = c.Form("form")
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
				query = c.Query("q")
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
				sock = c.Socket()
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

func TestEchoSupportResponseJSON(t *testing.T) {
	Convey("Given a echo.Server", t, func() {
		s := _echo.New()
		e := (s.Server()).(*echo.Echo)
		Convey(`A handler should be able respond using context.JSON`, func() {
			data := struct{ N int }{1}
			s.Get("/", func(c server.Context) error {
				return c.JSON(200, data)
			})
			_, body := request("GET", "/", e)
			Convey(`Response should be equal to {"N":1}`, func() {
				correct, _ := json.Marshal(data)
				So(body, ShouldEqual, string(correct)+"\n")
			})
		})
	})
}

func TestEchoSupportResponseXML(t *testing.T) {
	Convey("Given a echo.Server", t, func() {
		s := _echo.New()
		e := (s.Server()).(*echo.Echo)
		r := e.Router()
		Convey(`A handler should be able respond using context.XML`, func() {
			data := struct{ N int }{1}
			s.Get("/", func(c server.Context) error {
				return c.XML(200, data)
			})

			req, _ := http.NewRequest("GET", "/", nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			Convey(`Response ContentType should be equal to: `+echo.ApplicationXMLCharsetUTF8, func() {
				So(w.Header().Get(echo.ContentType), ShouldEqual, echo.ApplicationXMLCharsetUTF8)
			})
		})
	})
}

func TestEchoSupportResponseHTML(t *testing.T) {
	Convey("Given a echo.Server", t, func() {
		s := _echo.New()
		e := (s.Server()).(*echo.Echo)
		r := e.Router()
		Convey(`A handler should be able respond using context.HTML`, func() {
			data := struct{ N int }{1}
			s.Get("/", func(c server.Context) error {
				return c.HTML(200, "", data)
			})

			req, _ := http.NewRequest("GET", "/", nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			Convey(`Response code should be equal to: `+string(http.StatusOK), func() {
				So(w.Code, ShouldEqual, http.StatusOK)
			})
		})
	})
}

func TestEchoSupportResponseString(t *testing.T) {
	Convey("Given a echo.Server", t, func() {
		s := _echo.New()
		e := (s.Server()).(*echo.Echo)
		r := e.Router()
		Convey(`A handler should be able respond using context.String`, func() {
			data := struct{ N int }{1}
			s.Get("/", func(c server.Context) error {
				return c.String(200, "", data)
			})

			req, _ := http.NewRequest("GET", "/", nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			Convey(`Response code should be equal to: `+string(http.StatusOK), func() {
				So(w.Code, ShouldEqual, http.StatusOK)
			})
		})
	})
}

func TestEchoSupportResponseNoContent(t *testing.T) {
	Convey("Given a echo.Server", t, func() {
		s := _echo.New()
		e := (s.Server()).(*echo.Echo)
		r := e.Router()
		Convey(`A handler should be able respond using context.NoContent`, func() {
			s.Get("/", func(c server.Context) error {
				return c.NoContent(200)
			})

			req, _ := http.NewRequest("GET", "/", nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			Convey(`Response code should be equal to: `+string(http.StatusOK), func() {
				So(w.Code, ShouldEqual, http.StatusOK)
			})
		})
	})
}

func TestEchoSupportResponseError(t *testing.T) {
	Convey("Given a echo.Server", t, func() {
		s := _echo.New()
		e := (s.Server()).(*echo.Echo)
		Convey(`A handler should be able respond using context.Error(context, int, string)`, func() {
			s.Use(func(c server.Context) error {
				return c.Error(500, "error")
			})
			code, body := request("GET", "/", e)
			Convey(`Response status code should be 500`, func() {
				So(code, ShouldEqual, 500)
			})
			Convey(`Response body should be "error"`, func() {
				So(body, ShouldEqual, "error\n")
			})
		})

		Convey(`A handler should be able respond using context.Error(context, int, error)`, func() {
			s.Use(func(c server.Context) error {
				return c.Error(500, errors.New("error"))
			})
			code, body := request("GET", "/", e)
			Convey(`Response status code should be 500`, func() {
				So(code, ShouldEqual, 500)
			})
			Convey(`Response body should be "error"`, func() {
				So(body, ShouldEqual, "error\n")
			})
		})

		Convey(`A handler should be able respond using context.Error(context, int, any)`, func() {
			s.Use(func(c server.Context) error {
				return c.Error(500, 1)
			})
			code, _ := request("GET", "/", e)
			Convey(`Response status code should be 500`, func() {
				So(code, ShouldEqual, 500)
			})
		})

		Convey(`A handler should be able respond using context.Error(context, int)`, func() {
			s.Use(func(c server.Context) error {
				return c.Error(500)
			})
			code, _ := request("GET", "/", e)
			Convey(`Response status code should be 500`, func() {
				So(code, ShouldEqual, 500)
			})
		})
	})
}
