package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var Cfg *Config

type Config struct {
	es   Elasticsearch
	host string
	port int
}

func Load() error {
	if Cfg != nil {
		return nil
	}

	viper.AutomaticEnv()
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	viper.AddConfigPath("../")
	err := viper.ReadInConfig()
	if err != nil {
		_, ok := err.(viper.ConfigFileNotFoundError)
		if !ok {
			panic(err)
		}
	}

	Cfg = &Config{
		es:   newESConfig(),
		host: getStringOrPanic("APP_HOST"),
		port: getIntOrDefault("APP_PORT", 8080),
	}

	return nil
}

func (c *Config) Addr() string {
	return fmt.Sprintf("%s:%d", c.host, c.port)
}

func (c *Config) GetES() Elasticsearch {
	return c.es
}
