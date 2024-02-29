package infra

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

var ch *amqp.Channel

func channel() *amqp.Channel {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
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
