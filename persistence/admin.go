package persistence

import (
	"context"

	"dynamodbdemo/interfaces"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type StorageManager struct {
	DA interfaces.DynamoDbAdminAPI
}

func NewTodoTableManager(dyna interfaces.DynamoDbAdminAPI) *StorageManager {
	return &StorageManager{
		DA: dyna,
	}
}

func (s *StorageManager) CreateTodoTable() error {
	_, err := s.DA.CreateTable(context.TODO(), &dynamodb.CreateTableInput{
		TableName: aws.String("todo"),
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("id"),
				AttributeType: types.ScalarAttributeTypeN,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("id"),
				KeyType:       types.KeyTypeHash,
			},
		},
		BillingMode: types.BillingModePayPerRequest,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *StorageManager) DeleteTodoTable() error {
	_, err := s.DA.DeleteTable(context.TODO(), &dynamodb.DeleteTableInput{
		TableName: aws.String("todo"),
	})
	if err != nil {
		return err
	}
	return nil
}
