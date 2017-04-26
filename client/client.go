package client

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/paaff/PASO/config"
	"github.com/paaff/PASO/store"
	"github.com/paaff/PASO/wabbit"
)

// Start - Global function to start the client.
func Start(conf *config.Config) {
	// Start detection of bluetooth data
	dataChannel := make(chan store.BlueData)
	go detectBluetooth(dataChannel)

	w, err := wabbit.InitWabbitPublisher(conf.Username, conf.Pass, conf.Address, conf.Port, conf.ExchangeName, conf.ExchangeType, conf.RoutingKey)
	if err != nil {
		fmt.Println("Ini of publisher failed")
	}
	defer w.Connection.Close()
	defer w.Channel.Close()

	for data := range dataChannel {
		log.Printf("About to publish this phone: %s", data)
		jsonData, err := json.Marshal(data)
		if err != nil {
			fmt.Println("JSON Marshal went wrong")
		}
		if err = w.PublishMessage(jsonData, conf.ExchangeName, conf.RoutingKey); err != nil {
			fmt.Println("Publish went wrong.")
		}
	}

}
