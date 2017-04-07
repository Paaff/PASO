package server

import (
	"fmt"
	"log"
	"os"

	"github.com/paaff/PASO/config"
	"github.com/streadway/amqp"
)

// InitSRabbit is a function to be called by the servers main function enabling a connection to be made to the RabbitMQ server.
func InitSRabbit(conf *config.Config) <-chan amqp.Delivery {
	dialPath := fmt.Sprintf("amqp://%s:%s@%s:%s/", conf.Username, conf.Pass, conf.Address, conf.Port)
	conn, err := amqp.Dial(dialPath)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		conf.ExchangeName, // name
		conf.ExchangeType, // type
		true,              // durable
		false,             // auto-deleted
		false,             // internal
		false,             // no-wait
		nil,               // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	q, err := ch.QueueDeclare(
		"queuename", // name
		false,       // durable
		false,       // delete when usused
		true,        // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	failOnError(err, "Failed to declare a queue")

	if len(os.Args) < 2 {
		log.Printf("Usage: %s [binding_key]...", os.Args[0])
		os.Exit(0)
	}
	for _, s := range os.Args[1:] {
		log.Printf("Binding queue %s to exchange %s with routing key %s",
			q.Name, "logs_topic", s)
		err = ch.QueueBind(
			q.Name,       // queue name
			s,            // routing key
			"logs_topic", // exchange
			false,
			nil)
		failOnError(err, "Failed to bind a queue")
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	return msgs
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}