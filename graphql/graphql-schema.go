package main

import (
	"motif/graphql/resolvers/user"

	"github.com/graphql-go/graphql"
)

func GetQuery() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"getUser": user.GetGQL,
		},
	})
}
func GetMutation() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"createUser": user.CreateGQL,
		},
	})
}

func GetGraphQLSchema() (graphql.Schema, error) {
	rootQuery := GetQuery()
	rootMutation := GetMutation()
	return graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})
}
