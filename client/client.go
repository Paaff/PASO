package client

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/paaff/PASO/config"
	"github.com/paaff/PASO/rabbit"
	"github.com/streadway/amqp"
)

// Start - Global function to start the client.
func Start(conf *config.Config) {
	// Initialize rabbit connection and get ready to publish.
	dialPath := fmt.Sprintf("amqp://%s:%s@%s:%s/", conf.Username, conf.Pass, conf.Address, conf.Port)
	r, err := rabbit.NewRabbit(dialPath, conf.ExchangeName, conf.ExchangeType)
	if err != nil {
		log.Fatal(err)
	}

	// Start detection of bluetooth data
	dataChannel := make(chan blueData)
	go detectBluetooth(dataChannel)
	for data := range dataChannel {
		publish(data, r, conf)
	}

}

func publish(data blueData, r *rabbit.Rabbit, conf *config.Config) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	err = r.Channel.Publish(
		conf.ExchangeName, // exchange
		conf.RoutingKey,   // routing key
		false,             // mandatory
		false,             // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        jsonData,
		})

	if err != nil {
		log.Fatal(err)
	}

}
