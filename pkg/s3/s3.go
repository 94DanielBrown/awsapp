package s3

import (
	"log"

	"github.com/94DanielBrown/awsapp/pkg/awsconfig"
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








}
