package main

import (
	"github.com/erdnaxeli/ExLaSpatule/api/handlers"
	"github.com/labstack/echo/v4"
)

func runServer() {
	e := echo.New()
	c := newControllers()
	handlers.RegisterHandlers(e, c)

	e.File("/", "api/swagger-ui.html")
	e.File("/openapi.json", "api/openapi.json")

	e.Logger.Fatal(e.Start(":8123"))
}
