package echo_test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/backenderia/garf/server"
	_echo "github.com/backenderia/garf/server/echo"
	"github.com/labstack/echo"
	. "github.com/smartystreets/goconvey/convey"
)

func TestEchoSupportResponseJSON(t *testing.T) {
	Convey("Given a echo.Server", t, func() {
		s := _echo.New()
		e := (s.Server()).(*echo.Echo)
		Convey(`A handler should be able respond using server.JSON`, func() {
			data := struct{ N int }{1}
			s.Use(func(c server.Context) error {
				s.JSON(c, data)
				return nil
			})
			_, body := request("GET", "/", e)
			Convey(`Response should be equal to {"N":1}`, func() {
				correct, _ := json.Marshal(data)
				So(body, ShouldEqual, string(correct)+"\n")
			})
		})
	})
}

func TestEchoSupportResponseError(t *testing.T) {
	Convey("Given a echo.Server", t, func() {
		s := _echo.New()
		e := (s.Server()).(*echo.Echo)
		Convey(`A handler should be able respond using server.Error(context, int, string)`, func() {
			s.Use(func(c server.Context) error {
				return s.Error(c, 500, "error")
			})
			code, body := request("GET", "/", e)
			Convey(`Response status code should be 500`, func() {
				So(code, ShouldEqual, 500)
			})
			Convey(`Response body should be "error"`, func() {
				So(body, ShouldEqual, "error\n")
			})
		})

		Convey(`A handler should be able respond using server.Error(context, int, error)`, func() {
			s.Use(func(c server.Context) error {
				return s.Error(c, 500, errors.New("error"))
			})
			code, body := request("GET", "/", e)
			Convey(`Response status code should be 500`, func() {
				So(code, ShouldEqual, 500)
			})
			Convey(`Response body should be "error"`, func() {
				So(body, ShouldEqual, "error\n")
			})
		})

		Convey(`A handler should be able respond using server.Error(context, int, any)`, func() {
			s.Use(func(c server.Context) error {
				return s.Error(c, 500, 1)
			})
			code, _ := request("GET", "/", e)
			Convey(`Response status code should be 500`, func() {
				So(code, ShouldEqual, 500)
			})
		})

		Convey(`A handler should be able respond using server.Error(context, int)`, func() {
			s.Use(func(c server.Context) error {
				return s.Error(c, 500)
			})
			code, _ := request("GET", "/", e)
			Convey(`Response status code should be 500`, func() {
				So(code, ShouldEqual, 500)
			})
		})
	})
}
