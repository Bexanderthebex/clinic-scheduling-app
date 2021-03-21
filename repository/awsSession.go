package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	AWSSession "github.com/aws/aws-sdk-go/aws/session"
)

func Create(region string, profileName string) (*AWSSession.Session, error) {
	session, sessionCreationError := AWSSession.NewSession(
		&aws.Config{
			Region:      aws.String(region),
			Credentials: credentials.NewSharedCredentials("", profileName),
		})

	if sessionCreationError != nil {
		return session, sessionCreationError
	}

	return session, nil
}
