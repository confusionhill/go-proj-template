package hello

import (
	"github.com/labstack/echo/v4"
	"project-template/entity"
)

func (c HelloController) SayHello(ctx echo.Context) error {
	return ctx.JSON(200, entity.ResponseWrapper{
		Status: "200",
		Msg:    "Success",
		Data:   c.repo.GetHelloMessage(),
	})
}
