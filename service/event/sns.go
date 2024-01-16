package event

import (
	"encoding/json"
	"fmt"
	"image-receiver-api/config"
	"image-receiver-api/infra/aws"
	"image-receiver-api/schema"

	"github.com/aws/aws-sdk-go/service/sns"
)

var settings = config.GetSettings()

type EventSNS struct {
	SnsClient *sns.SNS
}

func NewEventSNS() *EventSNS {
	session := aws.GetSessionAWS()
	svc := sns.New(session)
	return &EventSNS{
		SnsClient: svc,
	}
}

func (s *EventSNS) Publish(message schema.Message) error {
	topicPtr := settings["TOPIC"]

	messageByte, err := json.Marshal(message)
	if err != nil {
		return err
	}
	messageString := string(messageByte)

	result, err := s.SnsClient.Publish(
		&sns.PublishInput{
			Message:  &messageString,
			TopicArn: &topicPtr,
		},
	)
	if err != nil {
		return nil
	}

	fmt.Println(*result.MessageId)
	return nil
}
