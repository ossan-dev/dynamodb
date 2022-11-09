package persistence

import (
	"context"
	"strconv"
	"time"

	"dynamodbdemo/interfaces"

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
	_, err := c.DC.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("todo"),
		Item: map[string]types.AttributeValue{
			"id":           &types.AttributeValueMemberN{Value: strconv.Itoa(id)},
			"category":     &types.AttributeValueMemberS{Value: category},
			"description":  &types.AttributeValueMemberS{Value: description},
			"is_completed": &types.AttributeValueMemberS{Value: "false"},
			"created_on":   &types.AttributeValueMemberS{Value: time.Now().Format("2006-01-02 15:04:05")},
		},
	})
	if err != nil {
		return err
	}
	return nil
}
