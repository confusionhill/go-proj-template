package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"project-template/internal/config"
)

func startServer(cfg *config.MainConfig, e *echo.Echo) error {
	addr := fmt.Sprintf(":%s", cfg.Server.Port)
	e.HideBanner = true
	return e.Start(addr)
}
