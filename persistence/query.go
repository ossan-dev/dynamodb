package persistence

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"

	"dynamodbdemo/interfaces"
	"dynamodbdemo/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type QueryManager struct {
	DQ interfaces.DynamoDbQueryAPI
}

func NewQueryManager(dynaClient interfaces.DynamoDbQueryAPI) *QueryManager {
	return &QueryManager{
		DQ: dynaClient,
	}
}

func (q *QueryManager) GetTodoById(id int) (*models.Todo, error) {
	var todo models.Todo
	var rawTodo string
	res, err := q.DQ.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String("todo"),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberN{Value: strconv.Itoa(id)},
		},
	})
	if err != nil {
		return nil, err
	}

	if err = attributevalue.Unmarshal(res.Item["todo"], &rawTodo); err != nil {
		return nil, err
	}

	if rawTodo == "" {
		return nil, errors.New("todo not found")
	}

	if err = json.Unmarshal([]byte(rawTodo), &todo); err != nil {
		return nil, err
	}

	return &todo, nil
}
