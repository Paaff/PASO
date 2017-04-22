package wabbit

import (
	"fmt"

	"github.com/NeowayLabs/wabbit"
	"github.com/NeowayLabs/wabbit/amqp"
)

// Wabbit struct
type Wabbit struct {
	Connection wabbit.Conn
	Channel    wabbit.Channel
}

// NewWabbit returns a Wabbit struct in which the connection and channel can be accessed.
func NewWabbit(username, pass, address, port, exchangeName, exchangeType string) (*Wabbit, error) {
	// Create wabbit instance
	w := &Wabbit{
		Connection: nil,
		Channel:    nil,
	}
	var err error

	// Dial connection
	dialPath := createDialPath(username, pass, address, port)
	w.Connection, err = amqp.Dial(dialPath)
	if err != nil {
		return nil, fmt.Errorf("Dial: %s", err)
	}

	// Get channel
	w.Channel, err = w.Connection.Channel()
	if err != nil {
		return nil, fmt.Errorf("Channel: %s", err)
	}

	// Declare exchange
	if err = w.Channel.ExchangeDeclare(
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
