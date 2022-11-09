package persistence

import (
	"context"
	"time"

	"dynamodbdemo/interfaces"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type WaitManager struct {
	DQ interfaces.DynamoDbWaiterAPI
}

func NewWaitManager(dynaWaiter interfaces.DynamoDbWaiterAPI) *WaitManager {
	return &WaitManager{
		DQ: dynaWaiter,
	}
}

func (q *WaitManager) WaitTodoTable() error {
	if err := q.DQ.Wait(context.TODO(), &dynamodb.DescribeTableInput{
		TableName: aws.String("todo"),
	}, 30*time.Second); err != nil {
		return err
	}
	return nil
}
