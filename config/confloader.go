package config

import (
	"encoding/json"
	"os"
)

// Config - This struct holds the fields for the general config given, whether it is server or client.
type Config struct {
	Type         string `json:"type"`
	Username     string `json:"username"`
	Pass         string `json:"password"`
	Address      string `json:"address"`
	Port         string `json:"port"`
	ExchangeName string `json:"exchangeName"`
	ExchangeType string `json:"exchangeType"`
	RoutingKey   string `json:"routingKey"`
}

// LoadConfig - Loads the config from the given conf path.
func LoadConfig(confPath string) (*Config, error) {
	var conf Config
	config, osErr := os.Open(confPath)
	if osErr != nil {
		return &conf, osErr
	}

	decoder := json.NewDecoder(config)
	if err := decoder.Decode(&conf); err != nil {
		return &conf, err
	}
	return &conf, nil

}

// Methods for config struct
