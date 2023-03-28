package main

import (
	"github.com/labstack/echo/v4"
	app "project-template/application"
	"project-template/internal/config"
)

func application(cfg *config.MainConfig) error {
	e := echo.New()
	rsc, err := app.InitResource(cfg)
	if err != nil {
		return err
	}
	r, err := app.InitRepo(rsc)
	if err != nil {
		return err
	}
	c := app.InitController(r)
	app.InitRouter(e, c)
	return startServer(cfg, e)
}
