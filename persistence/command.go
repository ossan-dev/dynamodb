package persistence

import (
	"context"
	"encoding/json"
	"strconv"

	"dynamodbdemo/interfaces"
	"dynamodbdemo/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type CommandManager struct {
	DC interfaces.DynamoDbCommandAPI
}

func NewCommandManager(dynaCmd interfaces.DynamoDbCommandAPI) *CommandManager {
	return &CommandManager{
		DC: dynaCmd,
	}
}

func (c *CommandManager) InsertTodo(id int, category, description string) error {
	todo := models.NewTodo(id, category, description)
	todoBytes, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	_, err = c.DC.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("todo"),
		Item: map[string]types.AttributeValue{
			"id":   &types.AttributeValueMemberN{Value: strconv.Itoa(id)},
			"todo": &types.AttributeValueMemberS{Value: string(todoBytes)},
		},
	})
	if err != nil {
		return err
	}
	return nil
}
