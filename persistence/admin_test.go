package persistence_test

import (
	"context"
	"errors"
	"testing"

	"dynamodbdemo/persistence"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/stretchr/testify/assert"
)

type mockDynamoDbAdminAPI struct {
	fnCreateTable func(ctx context.Context, params *dynamodb.CreateTableInput, optFns ...func(*dynamodb.Options)) (*dynamodb.CreateTableOutput, error)
}

func (m *mockDynamoDbAdminAPI) CreateTable(ctx context.Context, params *dynamodb.CreateTableInput, optFns ...func(*dynamodb.Options)) (*dynamodb.CreateTableOutput, error) {
	return m.fnCreateTable(ctx, params, optFns...)
}

func TestCreateTodoTable(t *testing.T) {
	testSuite := []struct {
		name          string
		fnCreateTable func(ctx context.Context, params *dynamodb.CreateTableInput, optFns ...func(*dynamodb.Options)) (*dynamodb.CreateTableOutput, error)
		errExpected   error
	}{
		{
			"TableCreated",
			func(ctx context.Context, params *dynamodb.CreateTableInput, optFns ...func(*dynamodb.Options)) (*dynamodb.CreateTableOutput, error) {
				return &dynamodb.CreateTableOutput{}, nil
			},
			nil,
		},
		{
			"TableNotCreated",
			func(ctx context.Context, params *dynamodb.CreateTableInput, optFns ...func(*dynamodb.Options)) (*dynamodb.CreateTableOutput, error) {
				return &dynamodb.CreateTableOutput{}, errors.New("issue on AWS")
			},
			errors.New("issue on AWS"),
		},
	}

	for _, tt := range testSuite {
		t.Run(tt.name, func(t *testing.T) {
			client := &mockDynamoDbAdminAPI{
				fnCreateTable: tt.fnCreateTable,
			}
			sut := persistence.NewTodoTableManager(client, nil)
			got := sut.CreateTodoTable()
			assert.Equal(t, got, tt.errExpected)
		})
	}
}
