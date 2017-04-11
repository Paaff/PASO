package client

import (
	"fmt"
	"log"

	"github.com/paaff/PASO/config"
	"github.com/paaff/PASO/rabbit"
	"github.com/streadway/amqp"
)

// InitCRabbit creates a connection to the local RabbitMQ server.
func InitCRabbit(conf *config.Config) {
	dialPath := fmt.Sprintf("amqp://%s:%s@%s:%s/", conf.Username, conf.Pass, conf.Address, conf.Port)
	rabbit, err := rabbit.NewRabbit(dialPath, conf.ExchangeName, conf.ExchangeType)

	//TODO: Should be the variable or something, check ContentType
	body := "detectedPhones"
	err = ch.Publish(
		conf.ExchangeName, // exchange
		conf.RoutingKey,   // routing key
		false,             // mandatory
		false,             // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	log.Printf(" [x] Sent %s", body)
}

func publishData() {

}
