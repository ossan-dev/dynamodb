package persistence_test

import (
	"context"
	"errors"
	"testing"

	"dynamodbdemo/models"
	"dynamodbdemo/persistence"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/stretchr/testify/assert"
)

type mockDynamoDbQueryAPI struct {
	fnGetItem func(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)
}

func (m *mockDynamoDbQueryAPI) GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error) {
	return m.fnGetItem(ctx, params, optFns...)
}

func TestGetTodoById(t *testing.T) {
	testSuite := []struct {
		name        string
		fnGetItem   func(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)
		todo        *models.Todo
		errExpected error
	}{
		{
			"TodoRetrieved",
			func(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error) {
				return &dynamodb.GetItemOutput{
					Item: map[string]types.AttributeValue{
						"todo": &types.AttributeValueMemberS{
							Value: `{"id":1,"category":"test","description":"test","is_completed":false}`,
						},
					},
				}, nil
			},
			&models.Todo{Id: 1, Category: "test", Description: "test", IsCompleted: false},
			nil,
		},
		{
			"AWSIssues",
			func(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error) {
				return &dynamodb.GetItemOutput{}, errors.New("issues on AWS")
			},
			nil,
			errors.New("issues on AWS"),
		},
		{
			"TodoNotFound",
			func(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error) {
				return &dynamodb.GetItemOutput{}, nil
			},
			nil,
			errors.New("todo not found"),
		},
	}

	for _, tt := range testSuite {
		t.Run(tt.name, func(t *testing.T) {
			client := &mockDynamoDbQueryAPI{
				fnGetItem: tt.fnGetItem,
			}
			sut := persistence.NewQueryManager(client)
			res, err := sut.GetTodoById(1)
			assert.Equal(t, tt.todo, res)
			assert.Equal(t, tt.errExpected, err)
		})
	}
}
