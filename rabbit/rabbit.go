package rabbit

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// Rabbit - "Abstract" rabbit struct
type Rabbit struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

// NewRabbit - Initialization function for a Rabbit struct
func NewRabbit(username string, pass string, address string, port string, exchangeName string, exchangeType string) (*Rabbit, error) {
	dialPath := createDialPath(username, pass, address, port)
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

func createDialPath(username string, pass string, address string, port string) string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s/", username, pass, address, port)
}
