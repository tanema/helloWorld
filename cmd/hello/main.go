package main

import (
	"context"
	"github.com/tanema/helloWorld/lib/netlify"
	"net/http"
)

func handler(ctx context.Context, req *http.Request) *netlify.Response {
	return netlify.Text("Hello Netlify")
}

func main() {
	netlify.Start(handler)
}
