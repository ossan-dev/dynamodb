package interfaces

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoDbTableCreateAPI interface {
	CreateTable(ctx context.Context, params *dynamodb.CreateTableInput, optFns ...func(*dynamodb.Options)) (*dynamodb.CreateTableOutput, error)
}

type DynamoDbTableDeleteAPI interface {
	DeleteTable(ctx context.Context, params *dynamodb.DeleteTableInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DeleteTableOutput, error)
}

type DynamoDbWaiterAPI interface {
	Wait(ctx context.Context, params *dynamodb.DescribeTableInput, maxWaitDur time.Duration, optFns ...func(*dynamodb.TableExistsWaiterOptions)) error
}

type DynamoDbQueryAPI interface {
	GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)
}

type DynamoDbCommandAPI interface {
	PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
}
