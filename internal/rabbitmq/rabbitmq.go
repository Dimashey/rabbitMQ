package rabbitmq

import "github.com/streadway/amqp"

type Service interface {
	Connect() error
}

type RabbitMQ struct{}

func (rmq *RabbitMQ) Connect() error {
	var err error

	_, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}

	return nil
}

func New() *RabbitMQ {
	return &RabbitMQ{}
}
