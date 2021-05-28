package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

// AWSCredential contains a credential keypair
type AWSCredential struct {
	accessKey string
	secretKey string
}

// Guesser verifies the existance of a specific access key and secret key combination
type Guesser struct {
	awsSession *session.Session
}

func (g *Guesser) verifyKey(credential *AWSCredential) (err error) {
	if g.awsSession == nil {
		g.awsSession, err = session.NewSession(aws.NewConfig().WithRegion(region))
		if err != nil {
			return
		}
	}
	g.awsSession.Config.Credentials = credentials.NewStaticCredentials(credential.accessKey, credential.secretKey, "")

	svc := sts.New(g.awsSession)
	_, err = svc.GetCallerIdentity(&sts.GetCallerIdentityInput{})
	return
}
