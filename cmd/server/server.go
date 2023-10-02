package main

import (
	"github.com/erdnaxeli/ExLaSpatule/api/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func runServer() {
	e := echo.New()
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
		LogLevel:  log.ERROR,
	}))

	e.File("/", "api/swagger-ui.html")
	e.File("/openapi.json", "api/openapi.json")

	c := newControllers()
	handlers.RegisterHandlers(e, c)

	e.Logger.Fatal(e.Start(":8123"))
}
