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

	fmt.Println("Go RabbitMQ start")

	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println("error setting up our application")
		fmt.Println(err)
	}
}
