package main

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type alias struct {
	channelId string
	members   []string
}

type store struct {
	client *dynamodb.DynamoDB
}

const tableName = "aliases"

func newStore() (store, error) {
	config := &aws.Config{
		Region:   aws.String("us-west-2"),
		Endpoint: aws.String("http://localhost:8000"),
	}

	sess := session.Must(session.NewSession(config))

	client := dynamodb.New(sess)

	err := createAliasTable(client)
	if err != nil {
		return store{}, err
	}

	return store{client}, nil
}

func (store *store) createAlias(channelId string) {
	// store.client.PutItem()
}

func createAliasTable(db *dynamodb.DynamoDB) error {
	input := dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("channelId"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("channelId"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(tableName),
	}

	_, err := db.CreateTable(&input)

	if err != nil {
		var resourceInUseErr *dynamodb.ResourceInUseException
		if !errors.As(err, &resourceInUseErr) {
			return fmt.Errorf("Error creating table: %w", err)
		}
	}

	return nil
}
