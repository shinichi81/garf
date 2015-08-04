package echo

import (
	"github.com/backenderia/garf/server"
	"github.com/labstack/echo"
)

// JSON forwards to echo.JSON
func (e *echoHandler) JSON(c server.Context, d interface{}) error {
	ec := c.(*echoContext)
	return ec.ctx.JSON(200, d)
}

// Error forwards to echo.Error
func (e *echoHandler) Error(c server.Context, code int, d ...interface{}) error {
	switch v := d[0].(type) {
	case error:
		return echo.NewHTTPError(code, v.Error())
	case string:
		return echo.NewHTTPError(code, v)
	default:
		return echo.NewHTTPError(code)
	}
}
