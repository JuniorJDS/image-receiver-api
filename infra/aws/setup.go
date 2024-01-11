package aws

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

var sess *session.Session

func Session() *session.Session {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-west-2"),
		Credentials: credentials.NewSharedCredentials("", "dev-go-account"),
	})
	if err != nil {
		log.Printf("Error to connect to aws: %s\n", err.Error())
	}

	return sess
}

func GetSessionAWS() *session.Session {
	if sess == nil {
		sess = Session()
	}
	return sess
}
