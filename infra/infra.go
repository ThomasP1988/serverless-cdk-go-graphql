package main

import (
	"errors"
	"motif/shared/config"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/constructs-go/constructs/v3"
	"github.com/aws/jsii-runtime-go"
)

type InfraStackProps struct {
	awscdk.StackProps
}

var stage config.Stage = config.DEV

func NewInfraStack(scope constructs.Construct, id string, props *InfraStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	rights := awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
		Effect: awsiam.Effect_ALLOW,
		Actions: &[]*string{
			jsii.String("execute-api:*"),
			jsii.String("dynamodb:*"),
		},
		Resources: &[]*string{
			jsii.String("*"),
		},
	})

	// The code that defines your stack goes here

	SetEndpoints(stack, &rights)
	SetUserTable(stack)

	return stack
}

func main() {

	// HandleCLIFlags()
	app := awscdk.NewApp(nil)
	configContext := app.Node().TryGetContext(jsii.String("stage"))
	if configContext != nil {
		stage = config.Stage(configContext.(string))
	}

	println("stage", stage)

	if stage != config.DEV && stage != config.STAGING && stage != config.PROD {
		panic(errors.New("unknown stage"))
	}

	config.GetConfig(&stage)

	NewInfraStack(app, "MotifAPIv2", &InfraStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	// return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	return &awscdk.Environment{
		//  Account: jsii.String("123456789012"),
		Region: jsii.String(config.Conf.Region),
	}

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
