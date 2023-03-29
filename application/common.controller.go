package application

import (
	"log"
	"project-template/internal/config"
	"project-template/internal/controller/auth"
	"project-template/internal/controller/hello"
)

type CommonController struct {
	helloController *hello.HelloController
	authController  *auth.AuthController
	cfg             *config.MainConfig
}

func InitController(r *CommonRepository, cfg *config.MainConfig) (*CommonController, error) {
	c := new(CommonController)
	c.cfg = cfg
	c.helloController = hello.Init(r.helloRepo)
	c.authController = auth.Init(r.authRepo)
	log.Println("CONTROLLER INTEGRATED")
	return c, nil
}
