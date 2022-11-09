package models

import (
	"fmt"
	"time"
)

type Todo struct {
	Id          int       `json:"id" dynamodbav:"id"` // partition key in DynamoDb
	Category    string    `json:"category" dynamodbav:"category"`
	Description string    `json:"description" dynamodbav:"description"`
	IsCompleted bool      `json:"is_completed" dynamodbav:"is_completed"`
	CreatedOn   time.Time `json:"created_on" dynamodbav:"created_on"`
}

func NewTodo(id int, category, description string) *Todo {
	return &Todo{
		Id:          id,
		Category:    category,
		Description: description,
		IsCompleted: false,
		CreatedOn:   time.Now(),
	}
}

func (t *Todo) PrintInfo() {
	fmt.Println("Todo details:")
	fmt.Printf("\tid: %d\n", t.Id)
	fmt.Printf("\tcategory: %q\n", t.Category)
	fmt.Printf("\tdescription: %q\n", t.Description)
	fmt.Printf("\tis_completed: %v\n", t.IsCompleted)
	fmt.Printf("\tcreated_on: %v\n", t.CreatedOn)
}
