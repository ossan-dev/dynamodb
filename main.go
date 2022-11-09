package main

import (
	"dynamodbdemo/models"
	"dynamodbdemo/persistence"
	"dynamodbdemo/utils"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var cfg *aws.Config

func main() {
	var err error
	if cfg, err = utils.GetAwsConfig(); err != nil {
		panic(err)
	}

	dynaClient := dynamodb.NewFromConfig(*cfg)
	dynaManager := persistence.NewTodoTableManager(dynaClient, dynaClient)
	dynaWaiter := dynamodb.NewTableExistsWaiter(dynaClient)
	waitManager := persistence.NewWaitManager(dynaWaiter)
	queryManager := persistence.NewQueryManager(dynaClient)
	cmdManager := persistence.NewCommandManager(dynaClient)

	// create table
	if err = dynaManager.CreateTodoTable(); err != nil {
		panic(err)
	}

	// deferred call for teardown logic
	defer func() {
		if err = dynaManager.DeleteTodoTable(); err != nil {
			panic(err)
		}
	}()

	// wait for the table creation
	if err = waitManager.WaitTodoTable(); err != nil {
		panic(err)
	}

	// write to DynamoDb
	if err = cmdManager.InsertTodo(1, "Programming", "Complete DynamoDb Tutorial"); err != nil {
		panic(err)
	}

	// get item by id
	var todo *models.Todo
	if todo, err = queryManager.GetTodoById(1); err != nil {
		panic(err)
	}

	todo.PrintInfo()
}
