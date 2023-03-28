package hello

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"project-template/entity"
)

func (r HelloRouter) SayHello(c echo.Context) error {
	res, err := r.controller.SayHello(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ResponseWrapper{
			Status: "400",
			Msg:    "Failed",
			Data:   nil,
		})
	}
	return c.JSON(200, entity.ResponseWrapper{
		Status: "200",
		Msg:    "Success",
		Data:   res,
	})
}
