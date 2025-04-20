package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port int
}

func NewConfigExample() *Config {
	return &Config{
		Port: 8080,
	}
}

func NewConfigFromYml() *Config {
	viper.AddConfigPath("/run/secrets")
	viper.SetConfigName("gin-template")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {

		return NewConfigExample()
	}
}
