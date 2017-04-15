package server

import (
	"fmt"
	"log"

	"github.com/paaff/PASO/config"
	"github.com/paaff/PASO/rabbit"
	"github.com/streadway/amqp"
)

// InitWorker is a function to be called by the servers main function enabling a connection to be made to the RabbitMQ server.
func initWorker(conf *config.Config) <-chan amqp.Delivery {
	dialPath := fmt.Sprintf("amqp://%s:%s@%s:%s/", conf.Username, conf.Pass, conf.Address, conf.Port)
	r, err := rabbit.NewRabbit(dialPath, conf.ExchangeName, conf.ExchangeType)
	if err != nil {
		log.Fatal(err)
	}

	queue, err := r.Channel.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Fatal(err)
	}

	err = r.Channel.QueueBind(
		queue.Name,        // queue name
		conf.RoutingKey,   // routing key
		conf.ExchangeName, // exchange
		false,
		nil)
	if err != nil {
		log.Fatal(err)
	}

	msgs, err := r.Channel.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto ack
		false,      // exclusive
		false,      // no local
		false,      // no wait
		nil,        // args
	)
	if err != nil {
		log.Fatal(err)
	}

	return msgs
}
