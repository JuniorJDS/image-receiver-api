package event

import (
	"flag"
	"fmt"
	"image-receiver-api/infra/aws"

	"github.com/aws/aws-sdk-go/service/sns"
)

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

func (s *EventSNS) Publish() error {
	msgPtr := flag.String("m", "", "The message to send to the subscribed users of the topic")
	topicPtr := flag.String("t", "", "The ARN of the topic to which the user subscribes")

	result, err := s.SnsClient.Publish(
		&sns.PublishInput{
			Message:  msgPtr,
			TopicArn: topicPtr,
		},
	)
	if err != nil {
		return nil
	}

	fmt.Println(*result.MessageId)
	return nil
}
