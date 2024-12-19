package main

import (
	"github.com/labstack/echo/v4"
	"github.com/yowger/pet-day-care-api-2/config"
)

func main() {
	e := echo.New()

	configPath := "."
	configName := ".env"
	config := config.LoadAppConfig(configPath, configName)

}
