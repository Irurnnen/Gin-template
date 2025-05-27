package config

import (
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// Config main struct config of this project
type Config struct {
	ServerConfig   *ServerConfig   `mapstructure:"server"`
	DatabaseConfig *DatabaseConfig `mapstructure:"database"`
	RedisConfig    *RedisConfig    `mapstructure:"redis"`
	LogLevel       string          `mapstructure:"log_level"`
	Debug          bool            `mapstructure:"debug"`
}

// ServerConfig struct config for http server
type ServerConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

// DatabaseConfig struct config for SQL database
type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}

// GetDSN get Domain Source Name from config
func (d *DatabaseConfig) GetDSN() string {
	DSN := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", d.User, d.Password, d.Host, d.Port, d.DBName)
	return DSN
}

type RedisConfig struct {
	Address     string        `mapstructure:"address"`
	Password    string        `mapstructure:"password"`
	User        string        `mapstructure:"user"`
	DB          int           `mapstructure:"db"`
	MaxRetries  int           `mapstructure:"max_retries"`
	DialTimeout time.Duration `mapstructure:"dial_timeout"`
	Timeout     time.Duration `mapstructure:"timeout"`
}

// NewConfigExample creates example config if marshal config exited with error
func NewConfigExample() *Config {
	return &Config{
		ServerConfig: &ServerConfig{
			Host: "0.0.0.0",
			Port: 8080,
		},
		DatabaseConfig: &DatabaseConfig{
			Host:     "hostname",
			Port:     5432,
			User:     "user",
			Password: "password",
			DBName:   "dbname",
		},
		LogLevel: "production",
	}
}

// NewConfig create new config
func NewConfig() *Config {
	viper.AddConfigPath("/run/secrets")
	viper.SetConfigName("gin-template")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Error().Err(err).Msg("Failed to read config, using example config")
		return NewConfigExample()
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Error().Err(err).Msg("Failed to unmarshal config, using example config")
		return NewConfigExample()
	}

	return &config
}

func NewConfigDebug() *Config {
	Config := NewConfig()
	Config.Debug = true // Set debug mode to true
	return Config       // Return the modified config
}
