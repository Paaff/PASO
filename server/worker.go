package server

import (
	"log"

	"github.com/paaff/PASO/config"
	"github.com/paaff/PASO/rabbit"
)

// InitWorker is a function to be called by the servers main function enabling a connection to be made to the RabbitMQ server.
func initWorker(conf *config.Config) {
	r := rabbit.NewRabbit(conf.Username, conf.Pass, conf.Address, conf.Port, conf.ExchangeName, conf.ExchangeType)
	r.Connection.Close()
	r.Channel.Close()

	queue, err := r.Channel.QueueDeclare(
		"testqueue", // name
		false,       // durable
		false,       // delete when usused
		true,        // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		log.Fatalf("Error in queue declare: %v", err)
	}

	err = r.Channel.QueueBind(
		queue.Name,        // queue name
		conf.RoutingKey,   // routing key
		conf.ExchangeName, // exchange
		false,
		nil)
	if err != nil {
		log.Fatalf("Error in Queuebind: %v", err)
	}

	msgs, err := r.Channel.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto ack
		false,      // exclusive
		false,      // no local
		false,      // no wait
		nil,        // args
	)
	if err != nil {
		log.Fatalf("Error in Channel consuming: %v", err)
	}

	go func() {
		for d := range msgs {
			log.Printf("Message recieved..: %s", string(d.Body))
			convertBTData(d)
		}
	}()
}
