package common

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func GetOne(client *dynamodb.Client, tableName *string, output interface{}, keys map[string]interface{}) (bool, error) {
	var keyCond *expression.KeyConditionBuilder

	for k, v := range keys {
		if keyCond == nil {
			newCond := expression.Key(k).Equal(expression.Value(v))
			keyCond = &newCond
		} else {
			newCond := expression.Key(k).Equal(expression.Value(v))
			keyCond.And(newCond)
		}
	}

	expr, err := expression.NewBuilder().WithKeyCondition(*keyCond).Build()
	if err != nil {
		return false, err
	}

	input := &dynamodb.QueryInput{
		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		TableName:                 tableName,
	}

	queryOutput, err := client.Query(context.TODO(), input)
	if err != nil {
		return false, err
	}

	if len(queryOutput.Items) == 0 {
		return true, nil
	}

	err = attributevalue.UnmarshalMap(queryOutput.Items[0], output)

	if err != nil {
		println("failed to unmarshal Items", err.Error())
		return false, err
	}

	return false, nil
}

func AddOne(client *dynamodb.Client, tableName *string, item interface{}) error {

	marshalledItem, err := attributevalue.MarshalMap(item)
	if err != nil {
		println("dynamodb AddOne, error marshalling item", err.Error())
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      marshalledItem,
		TableName: tableName,
	}

	putOutput, err := client.PutItem(context.TODO(), input)

	if err != nil {
		return err
	}

	attributevalue.UnmarshalMap(putOutput.Attributes, item)
	return nil
}
