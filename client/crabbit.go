package client

import (
	"fmt"
	"log"

	"github.com/paaff/PASO/config"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

// InitCRabbit creates a connection to the local RabbitMQ server.
func InitCRabbit(conf *config.Config) {
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

	//TODO: Should be the variable or something, check ContentType
	body := "detectedPhones"
	err = ch.Publish(
		conf.ExchangeName, // exchange
		conf.RoutingKey,   // routing key
		false,             // mandatory
		false,             // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", body)
}

func publishData() {

}
