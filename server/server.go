package server

import (
	"log"

	"github.com/paaff/server"
)

// Start - Global function to start the server.
func Start() {
	// Start connection with RabitMQ server.
	// TODO: Input config values.
	msgs := server.InitSRabbit()

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}

type sensorData struct {
	timestamp float64
}

type client struct {
	id string
	sensorData
}
