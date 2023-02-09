package connections

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type QueueClient struct {
	Connection *amqp.Connection
}

func NewQueueClient(user string, pass string, host string, port int) *QueueClient {
	Connection, err := amqp.Dial("amqp://user:password@localhost:5672/")
	if err != nil {
		log.Panic("Failed to connect to RabbitMQ")
	}
	return &QueueClient{
		Connection: Connection,
	}
}
