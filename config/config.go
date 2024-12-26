package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	PORT         string `mapstructure:"PORT"`
	DATABASE_URL string `mapstructure:"DATABASE_URL"`
}

func LoadAppConfig() *Config {
	var config *Config

	viper.SetConfigName(".")
	viper.AddConfigPath(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalf("Config file not found in path: %s\n", ".env")
		}

		log.Fatalf("Error reading config file: %v\n", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Error unmarshalling config: %v\n", err)
	}

	return config
}
