package main

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/awsapigatewayv2integrations"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/awss3assets"
	"github.com/aws/constructs-go/constructs/v3"
	"github.com/aws/jsii-runtime-go"
)

var lambdaEnv *map[string]*string = &map[string]*string{
	"stage": jsii.String(string(stage)),
}

func SetEndpoints(stack constructs.Construct, rights *awsiam.PolicyStatement) {
	var environment map[string]*string = map[string]*string{
		"CGO_ENABLED": jsii.String("0"),
		"GOOS":        jsii.String("linux"),
		"GOARCH":      jsii.String("amd64"),
	}

	lambdaGraphQL := awslambda.NewFunction(stack, SetName("GraphQL"), &awslambda.FunctionProps{
		Code: awslambda.NewAssetCode(jsii.String("../graphql"), &awss3assets.AssetOptions{
			Bundling: &awscdk.BundlingOptions{
				Image:       awslambda.Runtime_GO_1_X().BundlingDockerImage(),
				User:        jsii.String("root"),
				Environment: &environment,
				Command: &[]*string{
					jsii.String("bash"),
					jsii.String("-c"),
					jsii.String("go build -tags lambda -mod=vendor -o /asset-output/graphql"),
				},
			},
		}),
		Handler:       jsii.String("graphql"),
		Timeout:       awscdk.Duration_Seconds(jsii.Number(300)),
		Runtime:       awslambda.Runtime_GO_1_X(),
		Environment:   lambdaEnv,
		InitialPolicy: &[]awsiam.PolicyStatement{*rights},
	})

	integration := awsapigatewayv2integrations.NewLambdaProxyIntegration(&awsapigatewayv2integrations.LambdaProxyIntegrationProps{
		Handler:              lambdaGraphQL,
		PayloadFormatVersion: awsapigatewayv2.PayloadFormatVersion_VERSION_2_0(),
	})

	api := awsapigatewayv2.NewHttpApi(stack, SetName("HTTP-GraphQL"), &awsapigatewayv2.HttpApiProps{
		ApiName: SetName("http-graphql"),
	})

	api.AddRoutes(&awsapigatewayv2.AddRoutesOptions{
		Path:        jsii.String("/graphql/{operation}"),
		Integration: integration,
	})

	api.AddRoutes(&awsapigatewayv2.AddRoutesOptions{
		Path:        jsii.String("/graphql"),
		Integration: integration,
	})

	awscdk.NewCfnOutput(stack, jsii.String("GraphQL/Endpoint"), &awscdk.CfnOutputProps{
		Value: api.ApiEndpoint(),
	})

}
