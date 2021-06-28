package common

import (
	"context"
	"motif/shared/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	configAWS "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var AWSConfig *aws.Config

func SetAWSConfig(region *string, profile *string) error {
	var cfg aws.Config
	var err error

	if region == nil {
		conf := config.GetConfig(nil)
		region = &conf.Region
	}

	if profile == nil { // default profile, used in lambda functions
		cfg, err = configAWS.LoadDefaultConfig(context.TODO(), configAWS.WithRegion(*region))
	} else { // named profile, might be used for test or in local
		cfg, err = configAWS.LoadDefaultConfig(context.TODO(), configAWS.WithRegion(*region), configAWS.WithSharedConfigProfile(*profile))
	}

	if err != nil {
		println("unable to load SDK config, %v", err)
	}

	AWSConfig = &cfg

	return err
}

func GetDynamoDBClient() (*dynamodb.Client, error) {
	var err error
	if AWSConfig == nil {
		err = SetAWSConfig(nil, nil)
	}

	return dynamodb.NewFromConfig(*AWSConfig), err
}
