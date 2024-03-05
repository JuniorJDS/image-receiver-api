package infra

import (
	"image-receiver-api/config"

	amqp "github.com/rabbitmq/amqp091-go"
)

var ch *amqp.Channel
var settings = config.GetSettings()

func channel() *amqp.Channel {
	conn, err := amqp.Dial(settings["RABBITMQ_URL"])
	if err != nil {
		conn.Close()
		panic(1)
	}

	ch, err = conn.Channel()
	if err != nil {
		ch.Close()
		panic(1)
	}
	return ch
}

func OpenChannel() *amqp.Channel {
	if ch == nil {
		ch = channel()
	}
	return ch
}
