// +build lambda
package main

import (
	"context"
	"errors"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	println("stage", os.Getenv("stage"))
	schema, err := GetGraphQLSchema()

	if err != nil {
		println("err", err.Error())
		errorAPI := errors.New("error in the server GraphQL Schema")
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       errorAPI.Error(),
		}, errors.New("error in the server GraphQL Schema")
	}

	response, err := HandleGraphQL(req.Body, &schema)

	if err != nil {
		errorAPI := errors.New("a problem happened while handling your request")
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       errorAPI.Error(),
		}, errorAPI
	}
	println("string(response)", string(response))
	return events.APIGatewayProxyResponse{
		Body:       string(response),
		StatusCode: 200,
		Headers: map[string]string{
			"content-type": "application/json",
		},
	}, nil
}

func main() {

	lambda.Start(HandleRequest)
}
