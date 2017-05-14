package server

import (
	"fmt"
	"log"

	"github.com/paaff/PASO/config"
	"github.com/paaff/PASO/store"
	"github.com/paaff/PASO/wabbit"
	"github.com/paaff/PASO/web"
)

// Start - Global function to start the server.
func Start(conf *config.Config) {
	// Initialize the local DB of allowed clients.
	store.InitDB()

	// Start connection with RabitMQ server.
	w, err := wabbit.InitWabbitConsumer(conf.Username, conf.Pass, conf.Address, conf.Port, "bluetoothqueue", conf.ExchangeName, conf.ExchangeType, conf.RoutingKey)
	if err != nil {
		fmt.Println("Ini of consumer failed")
	}
	defer w.Connection.Close()
	defer w.Channel.Close()

	go w.ConsumeMessage("bluetoothqueue")
	web.InitBoard()

	// Make the server run forever with an unbuffered channel.
	forever := make(chan bool)

	log.Printf("Server is running. Press CTRL+C to exit.")
	<-forever
}
