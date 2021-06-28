package user

import (
	"errors"
	"motif/shared/repositories/user"

	"github.com/graphql-go/graphql"
)

const (
	CreateArgsName  = "name"
	CreateArgsEmail = "email"
)

var CreateArgs = graphql.FieldConfigArgument{
	CreateArgsName: &graphql.ArgumentConfig{
		Description: "name of the user",
		Type:        graphql.NewNonNull(graphql.String), //
	},
	CreateArgsEmail: &graphql.ArgumentConfig{
		Description: "email of the user",
		Type:        graphql.NewNonNull(graphql.String), //
	},
}

func Create(p graphql.ResolveParams) (interface{}, error) {

	userDB, err := user.GetInstance()

	if err != nil {
		println("error getting user db instance", err)
		return nil, errors.New("error while trying to get user")
	}

	newUser := &user.User{
		Name:  p.Args[CreateArgsName].(string),
		Email: p.Args[CreateArgsEmail].(string),
	}

	err = (*userDB).Add(newUser)

	if err != nil {
		println("error creating user", err.Error())
		return nil, errors.New("error while creating user")
	}

	return Reducer(newUser), nil
}

var CreateGQL = &graphql.Field{
	Args:    CreateArgs,
	Type:    Type,
	Resolve: Create,
}
