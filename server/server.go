package server

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/paaff/PASO/client"
	"github.com/paaff/PASO/config"
	"github.com/paaff/PASO/wabbit"
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
				fmt.Println("Unmarshalling went wrong")
			}
		}
	}()
	// Make the server run forever with an unbuffered channel.
	forever := make(chan bool)

	log.Printf("Server is running. Press CTRL+C to exit.")
	<-forever
}
