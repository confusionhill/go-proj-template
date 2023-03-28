package application

import (
	"log"
	"project-template/internal/controller/hello"
)

type CommonController struct {
	heloController *hello.HelloController
}

func InitController(r *CommonRepository) *CommonController {
	c := new(CommonController)
	c.heloController = hello.Init(r.helloRepo)
	log.Println("CONTROLLER INTEGRATED")
	return c
}
