package client

import (
	"log"

	"github.com/paaff/PASO/config"
	"github.com/paaff/PASO/rabbit"
	"github.com/streadway/amqp"
)

// Start - Global function to start the client.
func Start(conf *config.Config) {
	// Initialize rabbit connection and get ready to publish.
	r := rabbit.NewRabbit(conf.Username, conf.Pass, conf.Address, conf.Port, conf.ExchangeName, conf.ExchangeType)
	defer r.Connection.Close()
	defer r.Channel.Close()
	// Start detection of bluetooth data
	dataChannel := make(chan BlueData)
	go detectBluetooth(dataChannel)
	for data := range dataChannel {
		log.Printf("About to publish this phone: %s", data.Bdaddress)
		publish(data, r, conf)
	}

}

func publish(data BlueData, r *rabbit.Rabbit, conf *config.Config) {
	/*jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}*/

	err := r.Channel.Publish(
		conf.ExchangeName, // exchange
		conf.RoutingKey,   // routing key
		false,             // mandatory
		false,             // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("jsonData"),
		})

	if err != nil {
		log.Fatalf("It could not publish, err: %v", err)
	}

}
