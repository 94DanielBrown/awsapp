package awsapp

import (
	"context"
	"fmt"
	"time"

	"github.com/9danielbrown/awsapp/pkg/dynamo"
)

func InitDynamo(ctx context.Context, tableName string) (string, error) {
	client, err := dynamo.Connect()
	if err != nil {
		return "", fmt.Errorf("error connecting to dynamodb: %w", err)
	}

	exists, err := dynamo.Exists(ctx, client, tableName)
	if err != nil {
		return "", fmt.Errorf("error checking if DynamoDB table exists: %w", err)
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()

	if !exists {
		err = dynamo.Create(ctx, client, tableName)
		if err != nil {
			return "", fmt.Errorf("error creating dynamodb table: %w", err)
		}
		return "table created successfully", nil
	} else {
		return fmt.Sprintf("table %v already exists", tableName), nil
	}
}
