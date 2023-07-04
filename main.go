package main

import (
	"fmt"

	"github.com/Dimashey/rabbitMQ/internal/rabbitmq"
)

type App struct {
	Rmq *rabbitmq.RabbitMQ
}

func Run() error {
	rmq := rabbitmq.New()

	app := App{
		Rmq: rmq,
	}

	err := app.Rmq.Connect()
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer app.Rmq.Conn.Close()

	fmt.Println("Go RabbitMQ start")

	err = app.Rmq.Publish("Hi")

	if err != nil {
		return err
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println("error setting up our application")
		fmt.Println(err)
	}
}
