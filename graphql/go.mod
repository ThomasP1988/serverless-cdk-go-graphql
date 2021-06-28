module motif/graphql

go 1.16

require (
	github.com/aws/aws-lambda-go v1.24.0
	github.com/graphql-go/graphql v0.7.9
	motif/shared v0.0.1
)

replace motif/shared v0.0.1 => ./../shared
