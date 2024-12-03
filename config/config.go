package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AWSRegion    string
	AWSBucket    string
	SNSTopicARN  string
}

func LoadConfig() Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Failed to load .env: %v", err)
	}

	return Config{
		AWSRegion:   os.Getenv("AWS_REGION"),
		AWSBucket:   os.Getenv("AWS_BUCKET_NAME"),
		SNSTopicARN: os.Getenv("AWS_SNS_TOPIC_ARN"),
	}
}
