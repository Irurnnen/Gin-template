package config

import (
	"fmt"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Config struct {
	Server   Server   `yaml:"host"`
	Database Database `yaml:"database"`
	LogLevel string   `yaml:"log_level"`
	Debug    bool     `yaml:"debug"`
}

type Database struct {
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	User       string `yaml:"user"`
	Password   string `yaml:"password"`
	DBName     string `yaml:"db_name"`
	Secure     bool   `yaml:"secure"`
	DriverName string `yaml:"driver_name"`
}

func (d *Database) GetDSN() string {
	DSN := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", d.User, d.Password, d.Host, d.Port, d.DBName)
	if d.Secure {
		return DSN
	}
	return DSN + "?sslmode=disable"
}

type Server struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func NewConfigExample() *Config {
	return &Config{
		Server: Server{
			Host: "0.0.0.0",
			Port: 8080,
		},
		Database: Database{
			Host:     "hostname",
			Port:     5432,
			User:     "user",
			Password: "password",
			DBName:   "dbname",
		},
		LogLevel: "production",
	}
}

func NewConfig() *Config {
	viper.AddConfigPath("/run/secrets")
	viper.SetConfigName("gin-template")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		zap.L().Error("Failed to read config, using example config", zap.Error(err))
		return NewConfigExample()
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		zap.L().Error("Failed to parse config, using example config", zap.Error(err))
		return NewConfigExample()
	}

	return &config
}

func NewConfigDebug() *Config {
	Config := NewConfig()
	Config.Debug = true // Set debug mode to true
	return Config       // Return the modified config
}
