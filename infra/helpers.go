package main

import "github.com/aws/jsii-runtime-go"

func SetName(name string) *string {
	return jsii.String(string(stage) + "-" + name)
}
