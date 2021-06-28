package user

import (
	"motif/shared/config"
	common "motif/shared/repositories"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/uuid"
)

type Database struct {
	client    *dynamodb.Client
	tableName string
}

var databaseInstance *Database

func GetInstance() (*Database, error) {
	if databaseInstance == nil {
		err := New()
		if err != nil {
			return nil, err
		}
	}
	return databaseInstance, nil
}

func New() error {
	client, err := common.GetDynamoDBClient()

	databaseInstance = &Database{
		client:    client,
		tableName: config.Conf.Tables[config.USER].Name,
	}
	println("TableName user", databaseInstance.tableName)
	return err
}

func (udb *Database) GetOne(userID string) (*User, error) {
	user := &User{}
	doesntExist, err := common.GetOne(udb.client, &udb.tableName, user, map[string]interface{}{
		"userId": userID,
	})

	if doesntExist {
		return nil, err
	}

	return user, err
}

func (udb *Database) Add(newUser *User) error {
	newUser.UserID = uuid.New().String()
	newUser.DateCreated = time.Time{}
	return common.AddOne(udb.client, &udb.tableName, newUser)
}
