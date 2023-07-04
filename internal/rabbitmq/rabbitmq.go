package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

type Service interface {
	Connect() error
	Publish(message string) error
}

type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

func (r *RabbitMQ) Connect() error {
	var err error

	r.Conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}

	r.Channel, err = r.Conn.Channel()

	if err != nil {
		return err
	}

	_, err = r.Channel.QueueDeclare("TestQueue", false, false, false, false, nil)

	if err != nil {
		return err
	}

	return nil
}

func (r *RabbitMQ) Publish(message string) error {
	err := r.Channel.Publish("", "TestQueue", false, false, amqp.Publishing{ContentType: "text/plain", Body: []byte(message)})
	if err != nil {
		return err
	}

	fmt.Println("Successfully published message to a queue")
	return nil
}

func New() *RabbitMQ {
	return &RabbitMQ{}
}
