package persistence_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"dynamodbdemo/persistence"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/stretchr/testify/assert"
)

type mockDynamoDbWaiterAPI struct {
	fnWait func(ctx context.Context, params *dynamodb.DescribeTableInput, maxWaitDur time.Duration, optFns ...func(*dynamodb.TableExistsWaiterOptions)) error
}

func (m *mockDynamoDbWaiterAPI) Wait(ctx context.Context, params *dynamodb.DescribeTableInput, maxWaitDur time.Duration, optFns ...func(*dynamodb.TableExistsWaiterOptions)) error {
	return m.fnWait(ctx, params, maxWaitDur, optFns...)
}

func TestWaitTodoTable(t *testing.T) {
	testSuite := []struct {
		name        string
		fnWait      func(ctx context.Context, params *dynamodb.DescribeTableInput, maxWaitDur time.Duration, optFns ...func(*dynamodb.TableExistsWaiterOptions)) error
		errExpected error
	}{
		{
			"TableReady",
			func(ctx context.Context, params *dynamodb.DescribeTableInput, maxWaitDur time.Duration, optFns ...func(*dynamodb.TableExistsWaiterOptions)) error {
				return nil
			},
			nil,
		},
		{
			"AWSIssues",
			func(ctx context.Context, params *dynamodb.DescribeTableInput, maxWaitDur time.Duration, optFns ...func(*dynamodb.TableExistsWaiterOptions)) error {
				return errors.New("issues on AWS")
			},
			errors.New("issues on AWS"),
		},
	}

	for _, tt := range testSuite {
		t.Run(tt.name, func(t *testing.T) {
			client := &mockDynamoDbWaiterAPI{
				fnWait: tt.fnWait,
			}
			sut := persistence.NewWaitManager(client)
			got := sut.WaitTodoTable()
			assert.Equal(t, tt.errExpected, got)
		})
	}
}
