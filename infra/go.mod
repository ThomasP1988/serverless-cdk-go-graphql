module infra

go 1.16

require (
	github.com/aws/aws-cdk-go/awscdk v1.110.0-devpreview
	github.com/aws/constructs-go/constructs/v3 v3.3.87
	github.com/aws/jsii-runtime-go v1.30.0
	motif/shared v0.0.1
)

replace motif/shared v0.0.1 => ./../shared
