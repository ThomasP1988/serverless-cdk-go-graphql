# serverless-cdk-go-graphql
Boilerplate for graphql app in go hosted in a serverless way with CDK (infrastructure as code).

This boilerplate is fully written in Go, it uses aws-sdk-2 and graphql-go, you can develop in local and deploy on AWS to get a serverless endpoint.
I also included the workspace for VSCode.

# prerequisite
- AWS CDK
- Go > 1.16

# custom profile
If you want to deploy and test with a specific profile, you need to add profile property in cdk.json AND in shared/config/config

# testing in local
To test locally, go in the ./graphql folder and run "go run ."

# deploy
run "cdk deploy" in ./infra folder

if you need to deploy on staging or production:

for stage: 
- cdk deploy -c stage=staging

for prod: 
- cdk deploy -c stage=prod


