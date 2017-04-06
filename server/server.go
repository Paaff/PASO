package server

import (
	"log"

	"github.com/paaff/PASO/config"
)

// Start - Global function to start the server.
func Start(conf *config.Config) {
	// Start connection with RabitMQ server.
	// TODO: Input config values.
	msgs := InitSRabbit(conf)

	// Make the server run forever with an unbuffered channel.
	forever := make(chan bool)

	// TODO: Initialize worker to
	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}

type blueData struct {
	bdaddress string
	class     string
}

type client struct {
	id string
	blueData
}
