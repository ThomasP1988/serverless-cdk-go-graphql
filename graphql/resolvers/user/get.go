package user

import (
	"errors"
	"motif/shared/repositories/user"

	"github.com/graphql-go/graphql"
)

const (
	ArgsUserID = "userId"
)

var GetArgs = graphql.FieldConfigArgument{
	ArgsUserID: &graphql.ArgumentConfig{
		Description: "Id of the user requested",
		Type:        graphql.NewNonNull(graphql.String), //
	},
}

func Get(p graphql.ResolveParams) (interface{}, error) {

	userDB, err := user.GetInstance()

	if err != nil {
		println("error getting user db instance", err)
		return nil, errors.New("error while trying to get user")
	}

	userResult, err := (*userDB).GetOne(p.Args[ArgsUserID].(string))

	if err != nil {
		println("error getting user entity", err.Error())
		return nil, errors.New("error while trying to get user")
	}

	if userResult == nil {
		return nil, errors.New("user not found")
	}

	return Reducer(userResult), nil
}

var GetGQL = &graphql.Field{
	Args:    GetArgs,
	Type:    Type,
	Resolve: Get,
}
