package main

import (
	"learn-echo/configs"
	"learn-echo/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	configs.LoadEnv()
	configs.InitDatabase()

	machine := echo.New()

	routes.InitRoute(machine)

	machine.Start(":8181")
}
