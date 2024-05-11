package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	Host         string
	Port         string
	Username     string
	UserPassword string
	DbPassword   string
	DBName       string
	SSLMode      string
}

func (c *Config) InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
