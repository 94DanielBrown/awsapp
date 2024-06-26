package awsapp

import (
	"context"
	"fmt"
	"time"

	"github.com/94DanielBrown/awsapp/pkg/dynamo"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// InitDynamo creates a dynamodb table with specified name using AWS credentials from env variables.
func InitDynamo(ctx context.Context, tableName string) (*dynamodb.Client, string, error) {
	client, err := dynamo.Connect()
	if err != nil {
		return nil, "", fmt.Errorf("error connecting to dynamodb: %w", err)
	}

	exists, err := dynamo.Exists(ctx, client, tableName)
	if err != nil {
		return nil, "", fmt.Errorf("error checking if dynamodb table exists: %w", err)
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()

	if !exists {
		err = dynamo.Create(ctx, client, tableName)
		if err != nil {
			return nil, "", fmt.Errorf("error creating dynamodb table: %w", err)
		}
		return client, fmt.Sprintf("table %v created successfully", tableName), nil
	} else {
		return client, fmt.Sprintf("table %v already exists", tableName), nil
	}
}
