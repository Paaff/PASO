package config

import (
	"encoding/json"
	"fmt"
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
	configFile, osErr := os.Open(confPath)
	if osErr != nil {
		return &conf, osErr
	}
	return decodeConfig(configFile, &conf)
}

func decodeConfig(configFile *os.File, conf *Config) (*Config, error) {
	decoder := json.NewDecoder(configFile)
	if err := decoder.Decode(&conf); err != nil {
		return conf, fmt.Errorf("Error while decoding the configfile into a config struct, err: %v", err)
	}
	return conf, nil
}
