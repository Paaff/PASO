package server

import (
	"github.com/paaff/PASO/config"
	"github.com/paaff/PASO/rabbit"
)

// InitWorker is a function to be called by the servers main function enabling a connection to be made to the RabbitMQ server.
func initWorker(conf *config.Config) {
	r := rabbit.NewRabbit(conf.Username, conf.Pass, conf.Address, conf.Port, conf.ExchangeName, conf.ExchangeType)
	r.Consume(conf.RoutingKey, conf.ExchangeName)

}
