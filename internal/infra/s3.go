package infra

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"bookNotification/config"

	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Storage struct {
	Client *s3.Client
	Bucket string
	Key    string
}

func NewS3Storage(cfg *config.Config) (*S3Storage, error) {
	if cfg.AWSAccessKey == nil || cfg.AWSSecretKey == nil {
		return nil, fmt.Errorf("missing AWS credentials")
	}

	awsCfg, err := awsConfig.LoadDefaultConfig(context.TODO(),
		awsConfig.WithRegion(*cfg.AWSRegion),
		awsConfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(*cfg.AWSAccessKey, *cfg.AWSSecretKey, "")),
	)
	if err != nil {
		return nil, fmt.Errorf("error setting up AWS: %w", err)
	}

	return &S3Storage{
		Client: s3.NewFromConfig(awsCfg),
		Bucket: *cfg.AWSBucket,
		Key:    *cfg.AWSKey,
	}, nil
}

func (s *S3Storage) LoadPages(ctx context.Context) []int {
	output, err := s.Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: &s.Bucket,
		Key:    &s.Key,
	})
	if err != nil {
		log.Printf("File not found: %v", err)
		return []int{}
	}
	defer output.Body.Close()

	var pages []int
	if err := json.NewDecoder(output.Body).Decode(&pages); err != nil {
		log.Fatalf("Fail to decode pages from S3: %v", err)
	}
	return pages
}

func (s *S3Storage) SavePages(ctx context.Context, pages []int) {
	data, err := json.Marshal(pages)
	if err != nil {
		log.Fatalf("Fail to serialize JSON: %v", err)
	}
	_, err = s.Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: &s.Bucket,
		Key:    &s.Key,
		Body:   bytes.NewReader(data),
	})
	if err != nil {
		log.Fatalf("Fail to save page number in S3: %v", err)
	}
}
