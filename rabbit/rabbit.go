package rabbit

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// Rabbit - "Abstract" rabbit struct
type Rabbit struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

// NewRabbit - Initialization function for a Rabbit struct
func NewRabbit(username string, pass string, address string, port string, exchangeName string, exchangeType string) *Rabbit {
	dialPath := createDialPath(username, pass, address, port)
	conn := dialConnection(dialPath)
	ch := openChannel(conn)
	declareExchange(ch, exchangeName, exchangeType)
	r := Rabbit{conn, ch}
	return &r
}

func declareExchange(ch *amqp.Channel, exchangeName string, exchangeType string) {
	err := ch.ExchangeDeclare(
		exchangeName, // name
		exchangeType, // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		log.Fatal("Failed to declare an exchange")
	}
}

func openChannel(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Failed to open a channel")
	}
	return ch
}

func dialConnection(dialPath string) *amqp.Connection {
	conn, err := amqp.Dial(dialPath)
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ")
	}
	return conn
}

func createDialPath(username string, pass string, address string, port string) string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s/", username, pass, address, port)
}

// Consume - Declare's, binds and consumes
func (r *Rabbit) Consume(routingKey string, exchangeName string) {
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
		queue.Name,   // queue name
		routingKey,   // routing key
		exchangeName, // exchange
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
		}
	}()
}
