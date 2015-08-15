package echo_test

import (
	"testing"

	_echo "github.com/backenderia/garf/server/echo"
	"github.com/labstack/echo"
	. "github.com/smartystreets/goconvey/convey"
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
