package wabbit

import (
	"fmt"

	"github.com/NeowayLabs/wabbit"
	"github.com/NeowayLabs/wabbit/amqp"
)

// Wabbit struct
type Wabbit struct {
	connection wabbit.Conn
	channel    wabbit.Channel
}

// NewWabbit returns a Wabbit struct in which the connection and channel can be accessed.
func NewWabbit(username, pass, address, port, exchangeName, exchangeType string) (*Wabbit, error) {
	// Create wabbit instance
	w := &Wabbit{
		connection: nil,
		channel:    nil,
	}
	var err error

	// Dial connection
	dialPath := createDialPath(username, pass, address, port)
	w.connection, err = amqp.Dial(dialPath)
	if err != nil {
		return nil, fmt.Errorf("Dial: %s", err)
	}

	// Get channel
	w.channel, err = w.connection.Channel()
	if err != nil {
		return nil, fmt.Errorf("Channel: %s", err)
	}

	// Declare exchange
	if err = w.channel.ExchangeDeclare(
		exchangeName, // name of the exchange
		exchangeType, // type
		wabbit.Option{
			"durable":  true,
			"delete":   false,
			"internal": false,
			"noWait":   false,
		},
	); err != nil {
		return nil, fmt.Errorf("Exchange Declare: %s", err)
	}

	// Return the basic Wabbit
	return w, nil

}

func createDialPath(username string, pass string, address string, port string) string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s/", username, pass, address, port)
}
