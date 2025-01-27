package amqp

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func MustNew(url string) *amqp.Connection {
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatalf("failed to connect to amqp: %s", err)
	}

	return conn
}
