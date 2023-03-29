package application

import (
	"log"
	"project-template/internal/config"
	"project-template/internal/router/auth"
	"project-template/internal/router/hello"
)

type CommonRouter struct {
	HelloRouter *hello.HelloRouter
	AuthRouter  *auth.AuthRouter
}

func InitRouter(c *CommonController, cfg *config.MainConfig) (*CommonRouter, error) {
	r := new(CommonRouter)
	r.HelloRouter = hello.Init(c.helloController)
	r.AuthRouter = auth.Init(c.authController, cfg)

	log.Println("ROUTERS INTEGRATED")
	return r, nil
}
