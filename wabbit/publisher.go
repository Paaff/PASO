package wabbit

import (
	"fmt"

	"github.com/NeowayLabs/wabbit"
)

// InitWabbitPublisher creates a Wabbit and initializes it as a publisher
func InitWabbitPublisher(username, pass, address, port, exchangeName, exchangeType, routingKey string) (*Wabbit, error) {
	publisher, err := NewWabbit(username, pass, address, port, exchangeName, exchangeType)
	if err != nil {
		return nil, fmt.Errorf("Error in initializing Wabbit, error: %s", err)
	}
	return publisher, nil
}

// PublishMessage will publish the message.
func (w *Wabbit) PublishMessage(body []byte, exchangeName string, routingKey string) error {
	return w.Channel.Publish(
		exchangeName, // exchange
		routingKey,   // routing key
		body,
		wabbit.Option{
			"deliveryMode": 2,
			"contentType":  "application/json",
		})
}
