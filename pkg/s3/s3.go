package s3

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	s3types "github.com/aws/aws-sdk-go-v2/service/s3/types"

	"github.com/94DanielBrown/awsapp/pkg/awsconfig"
)

const (
	ContentTypeJPEG = "image/jpeg"
	ContentTypePNG  = "image/png"
)

type Client struct {
	S3 *s3.Client
}

// Connect initializes a new S3 client
func Connect() (*Client, error) {
	config, err := awsconfig.New()
	if err != nil {
		log.Fatal(err)
	}
	client := s3.NewFromConfig(config)
	return &Client{S3: client}, nil
}

// GeneratePresignedURL creates a pre-signed URL for uploading an image to S3
// The contentType parameter should be either image/jpeg or image/png
func (c *Client) GeneratePresignedURL(bucketName, key string, contentType string, expiry time.Duration) (string, error) {
	if contentType != ContentTypeJPEG && contentType != ContentTypePNG {
		return "", fmt.Errorf("invalid content type: %s", contentType)
	}
	presigner := s3.NewPresignClient(c.S3)
	acl := s3types.ObjectCannedACLBucketOwnerFullControl

	putObjectParams := &s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(key),
		ContentType: aws.String(contentType),
		ACL:         acl,
	}

	resp, err := presigner.PresignPutObject(context.Background(), putObjectParams, func(p *s3.PresignOptions) {
		p.Expires = expiry
	})
	if err != nil {
		return "", err
	}

	return resp.URL, nil
}
