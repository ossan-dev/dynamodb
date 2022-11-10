package persistence_test

import (
	"context"
	"errors"
	"testing"

	"dynamodbdemo/persistence"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/stretchr/testify/assert"
)

type mockDynamoDbTableCreateAPI struct {
	fnCreateTable func(ctx context.Context, params *dynamodb.CreateTableInput, optFns ...func(*dynamodb.Options)) (*dynamodb.CreateTableOutput, error)
}

func (m *mockDynamoDbTableCreateAPI) CreateTable(ctx context.Context, params *dynamodb.CreateTableInput, optFns ...func(*dynamodb.Options)) (*dynamodb.CreateTableOutput, error) {
	return m.fnCreateTable(ctx, params, optFns...)
}

type mockDynamoDbTableDeleteAPI struct {
	fnDeleteTable func(ctx context.Context, params *dynamodb.DeleteTableInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DeleteTableOutput, error)
}

func (m *mockDynamoDbTableDeleteAPI) DeleteTable(ctx context.Context, params *dynamodb.DeleteTableInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DeleteTableOutput, error) {
	return m.fnDeleteTable(ctx, params, optFns...)
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
			client := &mockDynamoDbTableCreateAPI{
				fnCreateTable: tt.fnCreateTable,
			}
			sut := persistence.NewTodoTableManager(client, nil)
			got := sut.CreateTodoTable()
			assert.Equal(t, tt.errExpected, got)
		})
	}
}

func TestDeleteTodoTable(t *testing.T) {
	testSuite := []struct {
		name          string
		fnDeleteTable func(ctx context.Context, params *dynamodb.DeleteTableInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DeleteTableOutput, error)
		errExpected   error
	}{
		{
			"TableDeleted",
			func(ctx context.Context, params *dynamodb.DeleteTableInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DeleteTableOutput, error) {
				return &dynamodb.DeleteTableOutput{}, nil
			},
			nil,
		},
		{
			"TableNotDeleted",
			func(ctx context.Context, params *dynamodb.DeleteTableInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DeleteTableOutput, error) {
				return &dynamodb.DeleteTableOutput{}, errors.New("issue on AWS")
			},
			errors.New("issue on AWS"),
		},
	}

	for _, tt := range testSuite {
		t.Run(tt.name, func(t *testing.T) {
			client := &mockDynamoDbTableDeleteAPI{
				fnDeleteTable: tt.fnDeleteTable,
			}
			sut := persistence.NewTodoTableManager(nil, client)
			got := sut.DeleteTodoTable()
			assert.Equal(t, tt.errExpected, got)
		})
	}
}
