package config

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
)

var Cfg *Config

type Config struct {
	es        Elasticsearch
	scenarios []Scenario
	host      string
	port      int
}

type Scenario struct {
	Case       string `json:"case"`
	Url        string `json:"url"`
	Duration   int    `json:"duration"`
	Rates      []int  `json:"rates"`
	MaxWorkers int    `json:"max_workers"`
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
		es:        newESConfig(),
		host:      getStringOrPanic("APP_HOST"),
		port:      getIntOrDefault("APP_PORT", 8080),
		scenarios: parseScenario(),
	}

	return nil
}

func parseScenario() (scenarios []Scenario) {
	rawString := getStringOrPanic("ES_SCENARIOS")
	err := json.Unmarshal([]byte(rawString), &scenarios)
	if err != nil {
		panic(err)
	}

	return
}

func (c *Config) Addr() string {
	return fmt.Sprintf("%s:%d", c.host, c.port)
}

func (c *Config) GetES() Elasticsearch {
	return c.es
}

func (c *Config) GetESScenarios() []Scenario {
	return c.scenarios
}
