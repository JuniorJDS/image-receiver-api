package event

import (
	"context"
	"encoding/json"
	infra "image-receiver-api/infra/rabbitMq"
	"image-receiver-api/schema"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type EventAmqp struct {
	Channel *amqp.Channel
}

func NewEventAmqp() *EventAmqp {
	return &EventAmqp{
		Channel: infra.OpenChannel(),
	}
}

func (e *EventAmqp) Publish(message schema.Message) error {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*8)
	defer cancel()

	msgByte, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = e.Channel.PublishWithContext(
		ctxTimeout,
		"go-images",
		"golang-api",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/json",
			Body:        msgByte,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

// func (e *EventAmqp) Consume(out chan amqp.Delivery, queue string) error {
// 	msgs, err := e.Channel.Consume(
// 		queue,
// 		"uploaded-image",
// 		false,
// 		false,
// 		false,
// 		false,
// 		nil,
// 	)
// 	if err != nil {
// 		return err
// 	}

// 	for msg := range msgs {
// 		out <- msg
// 	}

// 	return nil
// }
