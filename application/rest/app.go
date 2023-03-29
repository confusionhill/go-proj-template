package main

import (
	"github.com/go-playground/validator/v10"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	app "project-template/application"
	"project-template/internal/config"
	locMid "project-template/internal/middleware"
	"project-template/internal/router/health"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func application(cfg *config.MainConfig) error {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	rsc, err := app.InitResource(cfg)
	if err != nil {
		return err
	}
	r, err := app.InitRepo(rsc, cfg)
	if err != nil {
		return err
	}
	c, err := app.InitController(r, cfg)
	if err != nil {
		return err
	}
	router, err := app.InitRouter(c, cfg)
	if err != nil {
		return err
	}

	e.GET("/health", health.CheckHealth)
	v1 := e.Group("/v1")
	e.Use(middleware.CORSWithConfig(cfg.CorsCfg))
	e.Use(middleware.GzipWithConfig(cfg.GzipConf))

	v1.GET("/hello", router.HelloRouter.SayHello)
	v1.POST("/register", router.AuthRouter.RegisterUser)
	v1.POST("/login", router.AuthRouter.LoginUser)

	v1.GET("/", locMid.AuthGuard(locMid.TestAuthGuard), echojwt.WithConfig(cfg.JwtCfg))

	return startServer(cfg, e)
}
