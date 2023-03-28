package hello

import "github.com/labstack/echo/v4"

func (r HelloRouter) SayHello(c echo.Context) error {
	return r.controller.SayHello(c)
}
