package wabbit

import (
	"fmt"

	"github.com/NeowayLabs/wabbit"
)

// InitWabbitConsumer creates a Wabbit and initializes it as a consumer
func InitWabbitConsumer(username, pass, address, port, queueName, exchangeName, exchangeType, routingKey string) (<-chan wabbit.Delivery, error) {
	// Ini consumer wabbit
	consumer := &Wabbit{
		connection: nil,
		channel:    nil,
	}
	var err error
	consumer, err = NewWabbit(username, pass, address, port, exchangeName, exchangeType)
	if err != nil {
		return nil, fmt.Errorf("Error in initializing Wabbit, error: %s", err)
	}

	// Wabbit is running, extend it to be a consumer by declaring a queue
	queue, err := consumer.channel.QueueDeclare(
		queueName, // name of the queue
		wabbit.Option{
			"durable":   true,
			"delete":    false,
			"exclusive": false,
			"noWait":    false,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("Queue Declare: %s", err)
	}

	// Bind the queue to exchange using the routing key.
	if err = consumer.channel.QueueBind(
		queue.Name(), // name of the queue
		routingKey,   // bindingKey
		exchangeName, // sourceExchange
		wabbit.Option{
			"noWait": false,
		},
	); err != nil {
		return nil, fmt.Errorf("Queue Bind: %s", err)
	}

	// Begin to consume messages
	msgs, err := consumer.channel.Consume(
		queue.Name(), // name
		"",           // consumerTag,
		wabbit.Option{
			"noAck":     false,
			"exclusive": false,
			"noLocal":   false,
			"noWait":    false,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("Queue Consume: %s", err)
	}

	return msgs, nil

}
