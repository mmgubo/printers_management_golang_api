package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"mfundo.com/printers/entity"
)


type dynamoDBRepo struct {
	tableName string
}

func NewDynamoDBRepository() PrinterRepository {
	return &dynamoDBRepo{
		tableName: "printers",
	}
}
func createDynamoDBClient() *dynamodb.DynamoDB {
	// Create AWS Session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Return DynamoDB client
	return dynamodb.New(sess)
}

func (repo *dynamoDBRepo) Save(printer *entity.Printer) (*entity.Printer, error) {
	// Get a new DynamoDB client
	dynamoDBClient := createDynamoDBClient()

	// Transforms the post to map[string]*dynamodb.AttributeValue
	attributeValue, err := dynamodbattribute.MarshalMap(printer)
	if err != nil {
		return nil, err
	}

	// Create the Item Input
	item := &dynamodb.PutItemInput{
		Item:      attributeValue,
		TableName: aws.String(repo.tableName),
	}

	// Save the Item into DynamoDB
	_, err = dynamoDBClient.PutItem(item)
	if err != nil {
		return nil, err
	}

	return printer, err
}

func (repo *dynamoDBRepo) FindAll() ([]entity.Printer, error) {
	// Get a new DynamoDB client
	dynamoDBClient := createDynamoDBClient()

	// Build the query input parameters
	params := &dynamodb.ScanInput{
		TableName: aws.String(repo.tableName),
	}

	// Make the DynamoDB Query API call
	result, err := dynamoDBClient.Scan(params)
	if err != nil {
		return nil, err
	}
	var printers []entity.Printer = []entity.Printer{}
	for _, i := range result.Items {
		printer := entity.Printer{}

		err = dynamodbattribute.UnmarshalMap(i, &printer)

		if err != nil {
			panic(err)
		}
		printers = append(printers, printer)
	}
	return printers, nil
}

func (repo *dynamoDBRepo) FindByID(id string) (*entity.Printer, error) {
	// Get a new DynamoDB client
	dynamoDBClient := createDynamoDBClient()

	result, err := dynamoDBClient.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(repo.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				N: aws.String(id),
			},
		},
	})
	if err != nil {
		return nil, err
	}
	printer := entity.Printer{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &printer)
	if err != nil {
		panic(err)
	}
	return &printer, nil
}




// // Delete: TODO
// func (repo *dynamoDBRepo) Delete(printer *entity.Printer) error {
// 	return nil
// }


