package application

import (
	"log"
	"project-template/internal/repository/hello"
)

type CommonRepository struct {
	helloRepo *hello.HelloRepository
}

func InitRepo(rsc *CommonResource) (*CommonRepository, error) {
	r := new(CommonRepository)
	hRepo, err := hello.Init(rsc.Database, rsc.Redis, rsc.Mongo)
	if err != nil {
		return nil, err
	}
	r.helloRepo = hRepo
	log.Println("REPOSITORY INTEGRATED!")
	return r, nil
}
