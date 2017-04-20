package server

import (
	"encoding/json"
	"log"

	"github.com/paaff/PASO/config"
	"github.com/streadway/amqp"
)

// Start - Global function to start the server.
func Start(conf *config.Config) {
	// Start connection with RabitMQ server.
	msgs := initWorker(conf)

	// Make the server run forever with an unbuffered channel.
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Print("Message recieved..")
			convertBTData(d)
		}
	}()

	log.Printf("Server is running. Press CTRL+C to exit.")
	<-forever
}

func convertBTData(delivery amqp.Delivery) {
	var phone blueData

	err := json.Unmarshal(delivery.Body, &phone)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Phone address: %s\nPhone class: %s", phone.bdaddress, phone.class)

}

type blueData struct {
	bdaddress string
	class     string
}

type client struct {
	id string
	blueData
}
