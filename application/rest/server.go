package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"project-template/internal/config"
)

func startServer(cfg *config.MainConfig, e *echo.Echo) error {
	//e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	//	AllowOrigins: []string{"*", "localhost"},
	//	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	//}))
	//e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
	//	Level: 5,
	//}))
	//e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))
	//e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
	//	ErrorMessage: "Request Timeout!",
	//	Timeout:      30 * time.Second,
	//}))
	addr := fmt.Sprintf(":%s", cfg.Server.Port)
	return e.Start(addr)
}
