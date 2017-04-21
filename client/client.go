package client

import (
	"log"

	"github.com/paaff/PASO/config"
	"github.com/paaff/PASO/wabbit"
)

// Start - Global function to start the client.
func Start(conf *config.Config) {
	// Start detection of bluetooth data
	dataChannel := make(chan BlueData)
	go detectBluetooth(dataChannel)

	w, err := wabbit.InitWabbitPublisher(conf.Username, conf.Pass, conf.Address, conf.Port, conf.ExchangeName, conf.ExchangeType, conf.RoutingKey)
	if err != nil {
		log.Fatal(err)
	}
	for data := range dataChannel {
		log.Printf("About to publish this phone: %s", data.Bdaddress)
		w.PublishMessage(data.Bdaddress, conf.ExchangeName, conf.RoutingKey)
	}

}
