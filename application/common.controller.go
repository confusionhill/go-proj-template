package application

import (
	"log"
	"project-template/internal/controller/auth"
	"project-template/internal/controller/hello"
)

type CommonController struct {
	helloController *hello.HelloController
	authController  *auth.AuthController
}

func InitController(r *CommonRepository) (*CommonController, error) {
	c := new(CommonController)
	c.helloController = hello.Init(r.helloRepo)
	c.authController = auth.Init(r.authRepo)
	log.Println("CONTROLLER INTEGRATED")
	return c, nil
}
