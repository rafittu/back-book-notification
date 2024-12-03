package infra

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/joho/godotenv"
)

type SNSService struct {
	Client *sns.Client
	Topic  string
}

func NewSNSService() *SNSService {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Fail to load .env: %v", err)
	}

	accessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	region := os.Getenv("AWS_REGION")
	topicARN := os.Getenv("AWS_SNS_TOPIC_ARN")

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyID, secretAccessKey, "")),
	)
	if err != nil {
		log.Fatalf("Fail to setup AWS: %v", err)
	}

	return &SNSService{
		Client: sns.NewFromConfig(cfg),
		Topic:  topicARN,
	}
}

func (s *SNSService) PublishMessage(ctx context.Context, message string) {
	_, err := s.Client.Publish(ctx, &sns.PublishInput{
		Message:  &message,
		TopicArn: &s.Topic,
	})
	if err != nil {
		log.Fatalf("Fail to send SNS notification: %v", err)
	} else {
		log.Println("SNS notification sent successfully!")
	}
}
