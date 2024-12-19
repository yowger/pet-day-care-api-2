package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	DATABASE_URL string `mapstructure:"DATABASE_URL"`
}

func LoadAppConfig(configPath, configName string) *Config {
	var config *Config

	viper.SetConfigName(configName)
	viper.AddConfigPath(configPath)
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Printf("Config file not found in path: %s\n", configPath)
			os.Exit(2)
		}

		fmt.Printf("Error reading config file: %v\n", err)
		os.Exit(2)
	}

	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Error unmarshalling config: %v\n", err)
		os.Exit(2)
	}

	return config
}
