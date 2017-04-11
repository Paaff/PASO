package rabbit

import (
	"log"

	"github.com/streadway/amqp"
)

// Rabbit - "Abstract" rabbit struct
type Rabbit struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

// NewRabbit - Initialization function for a Rabbit struct
func NewRabbit(dialPath string, exchangeName string, exchangeType string) (*Rabbit, error) {
	conn, err := amqp.Dial(dialPath)
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ")
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Failed to open a channel")
		return nil, err
	}

	err = ch.ExchangeDeclare(
		exchangeName, // name
		exchangeType, // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		log.Fatal("Failed to declare an exchange")
		return nil, err
	}

	r := Rabbit{conn, ch}
	return &r, nil

}
