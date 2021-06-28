package user

import (
	"github.com/graphql-go/graphql"
)

var Type = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"userId": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"dateCreated": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
