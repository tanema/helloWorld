package main

import (
	"context"

	"github.com/tanema/helloWorld/lib/netlify"
)

func handler(ctx context.Context, req netlify.Request) *netlify.Response {
	return netlify.Text("Hello Netlify")
}

func main() {
	netlify.Start(handler)
}
