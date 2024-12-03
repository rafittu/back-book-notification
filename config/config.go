package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AWSAccessKey   *string
	AWSSecretKey   *string
	AWSRegion      *string
	AWSBucket      *string
	AWSKey         *string
	AWSSNSTopicARN *string
	TotalPages     int
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, fmt.Errorf("failed to load .env: %w", err)
	}

	totalPages, err := strconv.Atoi(os.Getenv("TOTAL_PAGES"))
	if err != nil {
		return nil, fmt.Errorf("invalid TOTAL_PAGES value: %w", err)
	}

	return &Config{
		AWSAccessKey:   strPtr(os.Getenv("AWS_ACCESS_KEY_ID")),
		AWSSecretKey:   strPtr(os.Getenv("AWS_SECRET_ACCESS_KEY")),
		AWSRegion:      strPtr(os.Getenv("AWS_REGION")),
		AWSBucket:      strPtr(os.Getenv("AWS_BUCKET_NAME")),
		AWSKey:         strPtr(os.Getenv("AWS_KEY")),
		AWSSNSTopicARN: strPtr(os.Getenv("AWS_SNS_TOPIC_ARN")),
		TotalPages:     totalPages,
	}, nil
}

func strPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
