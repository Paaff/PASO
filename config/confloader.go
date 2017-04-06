package config

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

// initLoad takes a configuration type "server" or "client" and returns the path respectively.
//TODO: Pretty weird function to have tbh.
func initLoad(confType string) (string, error) {
	var path string
	if confType == "server" {
		path = "./serverconf.json"
	} else if confType == "client" {
		path = "./clientconf.json"
	} else {
		return path, errors.New("Configuration type was not as expected")
	}
	return path, nil
}

// Config - This struct holds the fields for the general config given, whether it is server or client.
type Config struct {
	Username     string `json:"username"`
	Pass         string `json:"password"`
	Address      string `json:"address"`
	Port         string `json:"port"`
	ExchangeName string `json:"exchangeName"`
	ExchangeType string `json:"exchangeType"`
	RoutingKey   string `json:"routingKey"`
}

// LoadConfig - Loads the config from the given conf path.
func LoadConfig(confType string) *Config {
	confPath, pathErr := initLoad(confType)
	if pathErr != nil {
		log.Fatal(pathErr)
	}

	var conf Config
	config, osErr := os.Open(confPath)
	if osErr != nil {
		log.Fatal(osErr)
	}

	decoder := json.NewDecoder(config)
	if err := decoder.Decode(&conf); err != nil {
		log.Fatal(err)
	}
	return &conf

}

// Methods for config struct
