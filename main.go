package main

import (
	"github.com/yowger/pet-day-care-api-2/config"
)

func main() {
	configPath := "."
	configName := ".env"
	config := config.LoadAppConfig(configPath, configName)

}
