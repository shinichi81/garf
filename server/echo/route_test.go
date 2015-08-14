package echo_test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/backenderia/garf/server"
	_echo "github.com/backenderia/garf/server/echo"
	"github.com/labstack/echo"
	. "github.com/smartystreets/goconvey/convey"
)

func TestEchoRouteMethodSupport(t *testing.T) {
	Convey("Given a server.Echo", t, func() {
		s := _echo.New()
		e := (s.Server()).(*echo.Echo)
		r := e.Router()

		var routeTests = []struct {
			name   string
			method string
			router routeMethod
		}{
			{"echo.Get", "GET", s.Get},
			{"echo.Post", "POST", s.Post},
			{"echo.Put", "PUT", s.Put},
			{"echo.Delete", "DELETE", s.Del},
			{"echo.Patch", "PATCH", s.Patch},
			{"echo.Options", "OPTIONS", s.Options},
			{"echo.Head", "HEAD", s.Head},
		}

		for _, test := range routeTests {
			Convey("When adding a handler for the method "+test.name, func() {
				resp := []byte(test.name)
				q := ""
				handler := func(c server.Context) error {
					q = c.Request().URL.Query().Get("q")
					c.Response().WriteHeader(http.StatusOK)
					c.Response().Write(resp)
					return nil
				}

				test.router("/"+test.name, handler)
				req, _ := http.NewRequest(test.method, "/"+test.name+"?q=1", nil)
				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)

				Convey("The request query `q` should be '1'", func() {
					So(q, ShouldEqual, "1")
				})

				Convey("The reponse status should be 200", func() {
					So(w.Code, ShouldEqual, http.StatusOK)
				})
				Convey("The reponse body should be OK", func() {
					So(bytes.Equal(w.Body.Bytes(), resp), ShouldBeTrue)
				})
			})
		}

		for _, test := range routeTests {
			Convey("When adding a handler (using Handler()) for the method "+test.name, func() {
				resp := []byte(test.name)
				q := ""
				handler := func(c server.Context) error {
					q = c.Request().URL.Query().Get("q")
					c.Response().WriteHeader(http.StatusOK)
					c.Response().Write(resp)
					return nil
				}

				s.Handle(test.method, "/"+test.name, handler)
				req, _ := http.NewRequest(test.method, "/"+test.name+"?q=1", nil)
				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)

				Convey("The request query `q` should be '1'", func() {
					So(q, ShouldEqual, "1")
				})

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
		{"garf-default", func(c server.Context) error {
			passed = true
			return nil
		}},
		{"garf-default-wrap", server.HandlerFunc(func(c server.Context) error {
			passed = true
			return nil
		})},
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

func requestPost(method, path string, data io.Reader, e *echo.Echo) (int, string) {
	r, _ := http.NewRequest(method, path, data)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	e.ServeHTTP(w, r)
	return w.Code, w.Body.String()
}
