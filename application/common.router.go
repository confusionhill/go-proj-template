package application

import (
	"github.com/labstack/echo/v4"
	"log"
	"project-template/internal/router/hello"
)

type CommonRouter struct {
	helloRouter *hello.HelloRouter
}

func InitRouter(e *echo.Echo, c *CommonController) {
	r := new(CommonRouter)
	r.helloRouter = hello.Init(c.heloController)

	e.GET("/hello", r.helloRouter.SayHello)
	log.Println("ROUTERS INTEGRATED")
}
