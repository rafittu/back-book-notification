package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	AWSBucket      *string
	AWSKey         *string
	AWSSNSTopicARN *string
	TotalPages     int
}

func LoadConfig() (*Config, error) {
	totalPages, err := strconv.Atoi(os.Getenv("TOTAL_PAGES"))
	if err != nil {
		return nil, fmt.Errorf("invalid TOTAL_PAGES value: %w", err)
	}

	return &Config{
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
