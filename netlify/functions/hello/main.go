package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/tanema/helloWorld/lib/netlify"
)

func handler(request netlify.Request) (netlify.Response, error) {
	return netlify.Response{
		StatusCode: 200,
		Body:       "Hello Netlify",
	}, nil
}

func main() {
	lambda.Start(handler)
}
