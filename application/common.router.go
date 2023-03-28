package application

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"project-template/internal/router/auth"
	"project-template/internal/router/hello"
)

type CommonRouter struct {
	helloRouter *hello.HelloRouter
	authRouter  *auth.AuthRouter
}

func InitRouter(e *echo.Echo, c *CommonController) {
	r := new(CommonRouter)
	r.helloRouter = hello.Init(c.helloController)
	r.authRouter = auth.Init(c.authController)
	v1 := e.Group("/v1")
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Accept", "Content-Type", "X-Requested-With"},
		MaxAge:       0,
	}))

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	v1.GET("/hello", r.helloRouter.SayHello)
	v1.POST("/register", r.authRouter.RegisterNewUser)

	log.Println("ROUTERS INTEGRATED")
}
