package config

import (
	"log"
	"testing"
)

func TestClientConfigLoadedIntoStruct(t *testing.T) {
	path := "./clientconf_example.json"
	c, err := LoadConfig(path)
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
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