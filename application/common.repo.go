package application

import (
	"log"
	"project-template/internal/repository/auth"
	"project-template/internal/repository/hello"
)

type CommonRepository struct {
	helloRepo *hello.HelloRepository
	authRepo  *auth.AuthRepository
}

func InitRepo(rsc *CommonResource) (*CommonRepository, error) {
	r := new(CommonRepository)
	hRepo, err := hello.Init(rsc.Database, rsc.Redis, rsc.Mongo)
	if err != nil {
		return nil, err
	}
	r.helloRepo = hRepo
	aRepo, err := auth.Init(rsc.Database)
	if err != nil {
		return nil, err
	}
	r.authRepo = aRepo
	log.Println("REPOSITORY INTEGRATED!")
	return r, nil
}
