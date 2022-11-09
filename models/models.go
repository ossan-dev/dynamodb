package models

import "time"

type Todo struct {
	Id          int    // partition key in DynamoDb
	Category    string // sort key in DynamoDb
	Description string
	IsCompleted bool
	CreatedOn   time.Time
}
