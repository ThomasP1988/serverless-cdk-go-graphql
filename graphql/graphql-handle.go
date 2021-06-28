package main

import (
	"encoding/json"

	"github.com/graphql-go/graphql"
)

func HandleGraphQL(body string, schema *graphql.Schema) ([]byte, error) {
	var apollo map[string]interface{}
	var variables map[string]interface{}
	if err := json.Unmarshal([]byte(body), &apollo); err != nil {
		return nil, err
	}

	query := apollo["query"]
	if apollo["variables"] != nil {
		variables = apollo["variables"].(map[string]interface{})
	}

	result := graphql.Do(graphql.Params{
		Schema:         *schema,
		RequestString:  query.(string),
		VariableValues: variables,
	})
	return json.Marshal(result)

}
