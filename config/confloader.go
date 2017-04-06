package config

import (
	"encoding/json"
	"log"
	"os"
)

type config struct {
	Username     string `json:"username"`
	Pass         string `json:"password"`
	Address      string `json:"address"`
	Port         string `json:"port"`
	ExchangeName string `json:"exchangeName"`
	ExchangeType string `json:"exchangeType"`
	RoutingKey   string `json:"routingKey"`
}

func loadConfig(confName string) *config {
	var conf config
	config, osErr := os.Open(confName)
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
