package persistence

import (
	"context"
	"time"

	"dynamodbdemo/interfaces"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type QueryManager struct {
	DQ interfaces.DynamoDbQuery
}

func NewQueryManager(dynaWaiter interfaces.DynamoDbQuery) *QueryManager {
	return &QueryManager{
		DQ: dynaWaiter,
	}
}

func (q *QueryManager) WaitTodoTable() error {
	if err := q.DQ.Wait(context.TODO(), &dynamodb.DescribeTableInput{
		TableName: aws.String("todo"),
	}, 30*time.Second); err != nil {
		return err
	}
	return nil
}
