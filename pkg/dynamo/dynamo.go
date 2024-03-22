package dynamo

import (
	"context"
	"log"

	"github.com/94DanielBrown/awsapp/pkg/awsconfig"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// Connect to dynamodb
func Connect() (*dynamodb.Client, error) {
	config, err := awsconfig.New()
	if err != nil {
		log.Fatal(err)
	}
	client := dynamodb.NewFromConfig(config)
	return client, nil
}

// Create a dynamodb table
func Create(ctx context.Context, client *dynamodb.Client, tableName string) error {

	_, err := client.CreateTable(ctx, &dynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("id"),
				AttributeType: types.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("timestamp"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("id"),
				KeyType:       types.KeyTypeHash,
			},
			{
				AttributeName: aws.String("timestamp"),
				KeyType:       types.KeyTypeRange,
			},
		},
		TableName:   aws.String(tableName),
		BillingMode: types.BillingModePayPerRequest,
	})

	if err != nil {
		return err
	}

	return nil
}

// Exists checks if dynamodb table exists or not
func Exists(ctx context.Context, client *dynamodb.Client, tableName string) (bool, error) {
	p := dynamodb.NewListTablesPaginator(client, nil, func(o *dynamodb.ListTablesPaginatorOptions) {
		o.StopOnDuplicateToken = true
	})

	for p.HasMorePages() {
		out, err := p.NextPage(ctx)
		if err != nil {
			return false, err
		}

		for _, tn := range out.TableNames {
			if tn == tableName {
				return true, nil
			}
		}
	}
	return false, nil
}
