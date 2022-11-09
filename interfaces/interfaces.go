package interfaces

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoDbAdminAPI interface {
	CreateTable(ctx context.Context, params *dynamodb.CreateTableInput, optFns ...func(*dynamodb.Options)) (*dynamodb.CreateTableOutput, error)
	DeleteTable(ctx context.Context, params *dynamodb.DeleteTableInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DeleteTableOutput, error)
}

type DynamoDbQueryAPI interface {
	Wait(ctx context.Context, params *dynamodb.DescribeTableInput, maxWaitDur time.Duration, optFns ...func(*dynamodb.TableExistsWaiterOptions)) error
}

type DynamoDbCommandAPI interface {
	PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
}
