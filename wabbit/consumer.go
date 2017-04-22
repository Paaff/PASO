package wabbit

import (
	"fmt"

	"github.com/NeowayLabs/wabbit"
)

// InitWabbitConsumer creates a Wabbit and initializes it as a consumer
func InitWabbitConsumer(username, pass, address, port, queueName, exchangeName, exchangeType, routingKey string) (*Wabbit, error) {
	// Ini consumer wabbit
	consumer := &Wabbit{
		Connection: nil,
		Channel:    nil,
	}
	var err error
	consumer, err = NewWabbit(username, pass, address, port, exchangeName, exchangeType)
	if err != nil {
		return nil, fmt.Errorf("Error in initializing Wabbit, error: %s", err)
	}

	// Wabbit is running, extend it to be a consumer by declaring a queue
	queue, err := consumer.Channel.QueueDeclare(
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
	if err = consumer.Channel.QueueBind(
		queue.Name(), // name of the queue
		routingKey,   // bindingKey
		exchangeName, // sourceExchange
		wabbit.Option{
			"noWait": false,
		},
	); err != nil {
		return nil, fmt.Errorf("Queue Bind: %s", err)
	}

	return consumer, nil

}

// ConsumeMessage will consume send them to the channel given
func (w *Wabbit) ConsumeMessage(queueName string, recieverChan chan<- []byte) {
	// Begin to consume messages
	msgs, err := w.Channel.Consume(
		queueName, // name
		"",        // consumerTag,
		wabbit.Option{
			"autoAck":   true,
			"exclusive": false,
			"noLocal":   false,
			"noWait":    false,
		},
	)
	if err != nil {
		fmt.Printf("Queue Consume: %s", err)
	}
	for d := range msgs {
		recieverChan <- d.Body()
	}

}
