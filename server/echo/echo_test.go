package echo_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/backenderia/garf/server"
	_echo "github.com/backenderia/garf/server/echo"
	"github.com/labstack/echo"
	. "github.com/smartystreets/goconvey/convey"
)

type routeMethod func(string, server.HttpHandler)

func TestEchoRouteMethodSupport(t *testing.T) {
	Convey("Given a server.Echo", t, func() {
		s := _echo.New()
		e := (s.Server()).(*echo.Echo)
		r := e.Router()

		tests := []struct {
			name   string
			method string
			router routeMethod
		}{
			{"echo.Get", "GET", s.Get},
			{"echo.Post", "POST", s.Post},
			{"echo.Put", "PUT", s.Put},
			{"echo.Del", "DELETE", s.Del},
			{"echo.Patch", "PATCH", s.Patch},
			{"echo.Options", "OPTIONS", s.Options},
			{"echo.Head", "HEAD", s.Head},
		}

		for _, test := range tests {
			Convey("When adding a handler for the method "+test.name, func() {
				resp := []byte(test.name)
				handler := func(c server.Context) error {
					c.Response().WriteHeader(http.StatusOK)
					c.Response().Write(resp)
					return nil
				}

				test.router("/"+test.name, handler)
				req, _ := http.NewRequest(test.method, "/"+test.name, nil)
				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)

				Convey("The reponse status should be 200", func() {
					So(w.Code, ShouldEqual, http.StatusOK)
				})
				Convey("The reponse body should be OK", func() {
					So(bytes.Equal(w.Body.Bytes(), resp), ShouldBeTrue)
				})
			})
		}
	})
}

func TestEchoMiddlewareSupport(t *testing.T) {
	passed := false
	tests := []struct {
		name    string
		handler interface{}
	}{
		{"echo-default", func(c *echo.Context) error {
			passed = true
			return nil
		}},
		{"http-default", func(w http.ResponseWriter, h *http.Request) {
			passed = true
		}},
		{"echo-default-wrap", func(h echo.HandlerFunc) echo.HandlerFunc {
			return func(c *echo.Context) error {
				passed = true
				return nil
			}
		}},
	}

	for _, test := range tests {
		Convey("Given a server.Echo", t, func() {
			s := _echo.New()
			e := (s.Server()).(*echo.Echo)
			Convey("When adding the handler: "+test.name, func() {
				s.Use(test.handler)
				passed = false
				request("GET", "/", e)

				Convey("The flag passed should be TRUE", func() {
					So(passed, ShouldBeTrue)
				})
			})
		})
	}
}

func request(method, path string, e *echo.Echo) (int, string) {
	r, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	e.ServeHTTP(w, r)
	return w.Code, w.Body.String()
}
