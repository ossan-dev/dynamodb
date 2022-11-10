package persistence_test

import (
	"context"
	"errors"
	"testing"

	"dynamodbdemo/persistence"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/stretchr/testify/assert"
)

type mockDynamoDbCommandAPI struct {
	fnPutItem func(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
}

func (m *mockDynamoDbCommandAPI) PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error) {
	return m.fnPutItem(ctx, params, optFns...)
}

func TestInsertTodo(t *testing.T) {
	testSuite := []struct {
		name        string
		fnPutItem   func(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
		errExpected error
	}{
		{
			"TodoAdded",
			func(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error) {
				return &dynamodb.PutItemOutput{}, nil
			},
			nil,
		},
		{
			"AWSIssues",
			func(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error) {
				return &dynamodb.PutItemOutput{}, errors.New("AWS issues")
			},
			errors.New("AWS issues"),
		},
	}

	for _, tt := range testSuite {
		t.Run(tt.name, func(t *testing.T) {
			client := &mockDynamoDbCommandAPI{
				fnPutItem: tt.fnPutItem,
			}
			sut := persistence.NewCommandManager(client)
			got := sut.InsertTodo(1, "test", "test")
			assert.Equal(t, tt.errExpected, got)
		})
	}
}
