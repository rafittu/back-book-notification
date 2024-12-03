package infra

import (
	"context"
	"fmt"
	"log"

	"bookNotification/config"

	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

type SNSService struct {
	Client *sns.Client
	Topic  string
}

func NewSNSService(cfg *config.Config) (*SNSService, error) {
	if cfg.AWSSNSTopicARN == nil {
		return nil, fmt.Errorf("missing SNS topic ARN")
	}

	awsCfg, err := awsConfig.LoadDefaultConfig(context.TODO(),
		awsConfig.WithRegion(*cfg.AWSRegion),
		awsConfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(*cfg.AWSAccessKey, *cfg.AWSSecretKey, "")),
	)
	if err != nil {
		return nil, fmt.Errorf("error setting up AWS: %w", err)
	}

	return &SNSService{
		Client: sns.NewFromConfig(awsCfg),
		Topic:  *cfg.AWSSNSTopicARN,
	}, nil
}

func (s *SNSService) PublishMessage(ctx context.Context, message string) {
	_, err := s.Client.Publish(ctx, &sns.PublishInput{
		Message:  &message,
		TopicArn: &s.Topic,
	})
	if err != nil {
		log.Printf("fail to send SNS notification: %v", err)
	} else {
		log.Println("SNS notification sent successfully!")
	}
}
