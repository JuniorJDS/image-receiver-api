package event

import "image-receiver-api/schema"

type IEvent interface {
	Publish(message schema.Message) error
}
