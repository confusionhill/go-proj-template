package application

import (
	"log"
	"project-template/internal/config"
	"project-template/internal/repository/auth"
	"project-template/internal/repository/hello"
)

type CommonRepository struct {
	helloRepo *hello.HelloRepository
	authRepo  *auth.AuthRepository
	cfg       *config.MainConfig
}

func InitRepo(rsc *CommonResource, cfg *config.MainConfig) (*CommonRepository, error) {
	r := new(CommonRepository)
	r.cfg = cfg
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
