package netlify

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
)

type (
	Response struct {
		StatusCode        int                 `json:"statusCode"`
		Headers           map[string]string   `json:"headers"`
		MultiValueHeaders map[string][]string `json:"multiValueHeaders"`
		Body              string              `json:"body"`
		IsBase64Encoded   bool                `json:"isBase64Encoded,omitempty"`
	}
	Handler   func(context.Context, *http.Request) *Response
	BgHandler func(context.Context, *http.Request) error
)

func Start(handler Handler) {
	lambda.Start(func(ctx context.Context, req *http.Request) (*Response, error) {
		return handler(ctx, req), nil
	})
}

func StartBg(handler BgHandler)                       { lambda.Start(handler) }
func Resp(status int, payload string) *Response       { return &Response{StatusCode: status, Body: payload} }
func ErrStatus(status int, err error) *Response       { return Resp(status, err.Error()) }
func Err(err error) *Response                         { return ErrStatus(http.StatusInternalServerError, err) }
func Status(code int) *Response                       { return Resp(code, "") }
func OK() *Response                                   { return Status(http.StatusOK) }
func TextStatus(status int, payload string) *Response { return Resp(status, payload) }
func Text(payload string) *Response                   { return TextStatus(http.StatusOK, payload) }
func JSON(payload io.Reader) *Response                { return JSONStatus(http.StatusOK, payload) }

func JSONStatus(status int, payload io.Reader) *Response {
	data, err := json.Marshal(payload)
	if err != nil {
		return Err(fmt.Errorf("Could not marshal response body: %v", err))
	}
	return TextStatus(status, string(data))
}
