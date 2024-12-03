package infra

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
)

type S3Storage struct {
	Client *s3.Client
	Bucket string
	Key    string
}

func NewS3Storage() *S3Storage {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Fail to load .env: %v", err)
	}

	accessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	region := os.Getenv("AWS_REGION")
	bucket := os.Getenv("AWS_BUCKET_NAME")
	key := os.Getenv("AWS_KEY")

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyID, secretAccessKey, "")),
	)
	if err != nil {
		log.Fatalf("Error to setup AWS: %v", err)
	}

	return &S3Storage{
		Client: s3.NewFromConfig(cfg),
		Bucket: bucket,
		Key:    key,
	}
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
