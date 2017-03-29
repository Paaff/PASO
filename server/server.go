package server

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// Start - Global function to start the server.
func Start() {

	// Try to connect to the RabbitMQ server.
	// TODO: Use config file to state the username, password and address info
	conn, err := amqp.Dial("amqp://hubrabbit:pasopass@192.168.0.109:5672/")
	failOnsError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// Create a channel.
	ch, err := conn.Channel()
	failOnsError(err, "Failed to open a channel")
	defer ch.Close()

	// Declaring which queue the server should subscribe to
	// TODO: It should probably be told from the conf file which queues to subscribe to and how many.
	q, err := ch.QueueDeclare(
		"NameOfQueue", // name
		true,          // durable
		false,         // delete when usused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	failOnsError(err, "Failed to declare a queue")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	// We consume the messages from the queue asynchronously, which is why a go channel is used.

	// TODO: Auto-ack is false as we want to know if a message is lost, ie. the server died and the producer as to resend it.
	// Test this and document it.
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnsError(err, "Failed to register a consumer")

	// Create a channel with no buffer, that the main thread will wait on, such that it does not terminate.
	// No buffer channels are also called syncrhonous as reader will wait on writer (halting the execution of the program)
	forever := make(chan bool)

	// This goroutine will listen for messages in the msgs channel and print them.
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}

func failOnsError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

type sensorData struct {
	timestamp float64
}

type client struct {
	id string
	sensorData
}
