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
	w, err := wabbit.InitWabbitConsumer(conf.Username, conf.Pass, conf.Address, conf.Port, "bluetoothqueue", conf.ExchangeName, conf.ExchangeType, conf.RoutingKey)
	if err != nil {
		fmt.Println("Ini of consumer failed")
	}
	defer w.Connection.Close()
	defer w.Channel.Close()

	var btData client.BlueData
	serverChan := make(chan []byte)
	go w.ConsumeMessage("bluetoothqueue", serverChan)

	for jsonBlob := range serverChan {
		if err = json.Unmarshal(jsonBlob, &btData); err != nil {
			fmt.Println("Unmarshalling went wrong")
		}
		fmt.Println(btData.Bdaddress, btData.Class, "It worked ???")

	}

	// Make the server run forever with an unbuffered channel.
	forever := make(chan bool)

	log.Printf("Server is running. Press CTRL+C to exit.")
	<-forever
}
