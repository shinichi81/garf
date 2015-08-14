package echo_test

import (
	"testing"

	"github.com/backenderia/garf/server"
	_echo "github.com/backenderia/garf/server/echo"
	"github.com/labstack/echo"
	. "github.com/smartystreets/goconvey/convey"
)

func TestEchoSupportContext(t *testing.T) {
	Convey("Given a echo.Server", t, func() {
		s := _echo.New()
		e := (s.Server()).(*echo.Echo)
		Convey(`A handler should be able to context.Set("n", 123)`, func() {
			var n int
			s.Use(func(c server.Context) error {
				c.Set("n", 123)
				return nil
			})
			s.Use(func(c server.Context) error {
				n = (c.Get("n")).(int)
				return nil
			})
			request("GET", "/", e)
			Convey(`context.Get("n") should return 123`, func() {
				So(n, ShouldEqual, 123)
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
