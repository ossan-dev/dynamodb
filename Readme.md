# DynamoDb

Simple repo to show off some capabilities of AWS DynamoDb.  
Tools used for this repo:

- Go
- AWS sdk for Go version 2. You can download from [here](github.com/aws/aws-sdk-go-v2/service/dynamodb)
- Localstack  
  
In this repo, you can see how I usually implement tests on AWS sdk v2. To spin up the DynamoDb service issue `docker-compose up`, located at the `docker-compose.yml`'s location.

## CLI-Commands

- List tables: `aws --endpoint-url=http://localhost:4566 dynamodb list-tables`
- Delete table: `aws --endpoint-url=http://localhost:4566 dynamodb delete-table --table-name <your table name>`
- List DynamoDb items: `aws --endpoint-url=http://localhost:4566 dynamodb scan --table-name <your table name>`
