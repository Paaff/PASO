package config

import (
	"log"
	"os"
	"testing"
)

func TestClientConfigLoadedIntoStruct(t *testing.T) {
	path := "./clientconf_example.json"
	c, err := LoadConfig(path)
	if err != nil {
		t.Errorf("An error appeared, err:%v", err)
	}
	if c.Type != "CLIENT" {
		t.Errorf("Not correct type: %v", c.Type)
	}

	if c.Username != "USERNAME" {
		t.Errorf("Not correct username: %v", c.Username)
	}

	if c.Pass != "PASSWORD" {
		t.Errorf("Not correct password: %v", c.Pass)
	}

	if c.Address != "ADDRESS" {
		t.Errorf("Not correct address: %v", c.Address)
	}

	if c.Port != "PORT" {
		t.Errorf("Not correct port: %v", c.Port)
	}

	if c.ExchangeName != "EXCHANGE_NAME" {
		t.Errorf("Not correct exchange name: %v", c.ExchangeName)
	}

	if c.ExchangeType != "EXCHANGE_TYPE" {
		t.Errorf("Not correct exchane type: %v", c.ExchangeType)
	}

	if c.RoutingKey != "ROUTING_KEY" {
		t.Errorf("Not correct routing key: %v", c.RoutingKey)
	}

}

func TestServerConfigLoadedIntoStruct(t *testing.T) {
	path := "./serverconf_example.json"
	c, err := LoadConfig(path)
	if err != nil {
		t.Errorf("An error appeared, err:%v", err)
	}

	if c.Type != "SERVER" {
		t.Errorf("Not correct type: %v", c.Type)
	}

	if c.Username != "USERNAME" {
		t.Errorf("Not correct username: %v", c.Username)
	}

	if c.Pass != "PASSWORD" {
		t.Errorf("Not correct password: %v", c.Pass)
	}

	if c.Address != "ADDRESS" {
		t.Errorf("Not correct address: %v", c.Address)
	}

	if c.Port != "PORT" {
		t.Errorf("Not correct port: %v", c.Port)
	}

	if c.ExchangeName != "EXCHANGE_NAME" {
		t.Errorf("Not correct exchange name: %v", c.ExchangeName)
	}

	if c.ExchangeType != "EXCHANGE_TYPE" {
		t.Errorf("Not correct exchane type: %v", c.ExchangeType)
	}

	if c.RoutingKey != "ROUTING_KEY" {
		t.Errorf("Not correct routing key: %v", c.RoutingKey)
	}
}

func TestWrongPathConfig(t *testing.T) {
	wrongPath := "wrongpath"
	_, err := LoadConfig(wrongPath)

	if err == nil {
		t.Errorf("Error was nil, but was expected to be not nil")
	}
}

func TestWrongConfigJSONLayout(t *testing.T) {
	file, err := os.Open("wrongJsonContent.json")
	if err != nil {
		log.Fatalf("Error in the creation of a testing mock file, err: %v", err)
	}

	var conf Config
	_, err = decodeConfig(file, &conf)
	if err == nil {
		if conf.Type != "TYPE" || conf.Username != "USERNAME" || conf.Pass != "PASSWORD" ||
			conf.Address != "ADDRESS" || conf.Port != "PORT" || conf.ExchangeName != "EXCHANGE_NAME" ||
			conf.ExchangeType != "EXCHANGE_TYPE" || conf.RoutingKey != "ROUTING_KEY" {
			t.Errorf("The config struct was not loaded properly, config is: %v", conf)
		}
	}
}

func TestWrongJSONFormat(t *testing.T) {
	file, err := os.Open("wrongJSONFormat.json")
	if err != nil {
		log.Fatalf("Error in the creation of a testing mock file, err: %v", err)
	}

	var conf Config
	_, err = decodeConfig(file, &conf)
	if err != nil {

	} else {
		t.Error(err)
		if conf.Type != "TYPE" || conf.Username != "USERNAME" || conf.Pass != "PASSWORD" ||
			conf.Address != "ADDRESS" || conf.Port != "PORT" || conf.ExchangeName != "EXCHANGE_NAME" ||
			conf.ExchangeType != "EXCHANGE_TYPE" || conf.RoutingKey != "ROUTING_KEY" {
			t.Errorf("The config struct was not loaded properly, config is: %v", conf)
		}
	}
}
