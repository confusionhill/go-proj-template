package hello

import hr "project-template/internal/repository/hello"

type HelloController struct {
	repo *hr.HelloRepository
}

func Init(r *hr.HelloRepository) *HelloController {
	c := &HelloController{repo: r}
	return c
}
