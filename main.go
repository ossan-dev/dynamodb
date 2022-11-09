package main

import (
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
	dynaManager := persistence.NewTodoTableManager(dynaClient)
	dynaWaiter := dynamodb.NewTableExistsWaiter(dynaClient)
	queryManager := persistence.NewQueryManager(dynaWaiter)
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
	if err = queryManager.WaitTodoTable(); err != nil {
		panic(err)
	}

	// write to DynamoDb
	if err = cmdManager.InsertTodo("Programming", "Complete DynamoDb Tutorial"); err != nil {
		panic(err)
	}
}
