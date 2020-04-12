package main

import (
	"os"

	"github.com/streadway/amqp"
)

func main() {
	amqpURL := os.Getenv("AMQP_URL")
	if amqpURL == "" {
		amqpURL = "amqp://guest:guest@localhost:5672"
	}

	connection, err := amqp.Dial(amqpURL)
	if err != nil {
		panic("could not establish AMQP connection: " + err.Error())
	}
	defer connection.Close()

	channel, err := connection.Channel()
	if err != nil {
		panic("could not open channel: " + err.Error())
	}
	//defer channel.Close()

	err = channel.ExchangeDeclare("events", "topic", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	message := amqp.Publishing{
		Body: []byte("Hello World"),
	}
	err = channel.Publish("events", "some-routing-key", false, false, message)
	if err != nil {
		panic("error while publishing message: " + err.Error())
	}
}
