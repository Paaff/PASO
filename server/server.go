package server

import (
	"encoding/json"
	"log"

	"github.com/paaff/PASO/client"
	"github.com/paaff/PASO/config"
	"github.com/paaff/PASO/wabbit"
	"github.com/streadway/amqp"
)

// Start - Global function to start the server.
func Start(conf *config.Config) {
	// Start connection with RabitMQ server.
	msgs, err := wabbit.InitWabbitConsumer(conf.Username, conf.Pass, conf.Address, conf.Port, "bluetoothqueue", conf.ExchangeName, conf.ExchangeType, conf.RoutingKey)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for d := range msgs {
			data := &client.BlueData{
				Bdaddress: "",
				Class:     "",
			}
			err := json.Unmarshal(d.Body(), &data)
			if err != nil {
				log.Fatal(err)
			}
		}
	}()
	// Make the server run forever with an unbuffered channel.
	forever := make(chan bool)

	log.Printf("Server is running. Press CTRL+C to exit.")
	<-forever
}

func convertBTData(delivery amqp.Delivery) {
	var phone client.BlueData

	err := json.Unmarshal(delivery.Body, &phone)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Phone address: %s\nPhone class: %s", phone.Bdaddress, phone.Class)

}
