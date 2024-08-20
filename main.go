package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context) (*string, error) {
	resp := "hello world"
	return &resp, nil
}

func main() {
	lambda.Start(HandleRequest)
}
