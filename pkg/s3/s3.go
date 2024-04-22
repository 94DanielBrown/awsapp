package s3

import (
	"context"
	"log"
	"time"

	"github.com/94DanielBrown/awsapp/pkg/awsconfig"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func Connect() (*s3.Client, error) {
	config, err := awsconfig.New()
	if err != nil {
		log.Fatal(err)
	}
	client := s3.NewFromConfig(config)
	return client, nil
}

// GeneratePresignedURL creates a pre-signed URL for uploading an image to S3
func GeneratePresignedURL(client *s3.Client, bucketName, key string, expiry time.Duration) (string, error) {
	presigner := s3.NewPresignClient(client)

	putObjectParams := &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	}

	resp, err := presigner.PresignPutObject(context.Background(), putObjectParams, func(p *s3.PresignOptions) {
		p.Expires = expiry
	})
	if err != nil {
		return "", err
	}

	return resp.URL, nil
}
