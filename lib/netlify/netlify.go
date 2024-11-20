package netlify

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
)

type (
	Request struct {
		Resource                        string              `json:"resource"` // The resource path defined in API Gateway
		Path                            string              `json:"path"`     // The url path for the caller
		HTTPMethod                      string              `json:"httpMethod"`
		Headers                         map[string]string   `json:"headers"`
		MultiValueHeaders               map[string][]string `json:"multiValueHeaders"`
		QueryStringParameters           map[string]string   `json:"queryStringParameters"`
		MultiValueQueryStringParameters map[string][]string `json:"multiValueQueryStringParameters"`
		PathParameters                  map[string]string   `json:"pathParameters"`
		StageVariables                  map[string]string   `json:"stageVariables"`
		RequestContext                  Context             `json:"requestContext"`
		Body                            string              `json:"body"`
		IsBase64Encoded                 bool                `json:"isBase64Encoded,omitempty"`
	}
	Context struct {
		AccountID         string                 `json:"accountId"`
		ResourceID        string                 `json:"resourceId"`
		OperationName     string                 `json:"operationName,omitempty"`
		Stage             string                 `json:"stage"`
		DomainName        string                 `json:"domainName"`
		DomainPrefix      string                 `json:"domainPrefix"`
		RequestID         string                 `json:"requestId"`
		ExtendedRequestID string                 `json:"extendedRequestId"`
		Protocol          string                 `json:"protocol"`
		Identity          Identity               `json:"identity"`
		ResourcePath      string                 `json:"resourcePath"`
		Path              string                 `json:"path"`
		Authorizer        map[string]interface{} `json:"authorizer"`
		HTTPMethod        string                 `json:"httpMethod"`
		RequestTime       string                 `json:"requestTime"`
		RequestTimeEpoch  int64                  `json:"requestTimeEpoch"`
		APIID             string                 `json:"apiId"` // The API Gateway rest API Id
	}
	Identity struct {
		CognitoIdentityID             string `json:"cognitoIdentityId,omitempty"`
		CognitoIdentityPoolID         string `json:"cognitoIdentityPoolId,omitempty"`
		AccountID                     string `json:"accountId,omitempty"`
		Caller                        string `json:"caller,omitempty"`
		APIKey                        string `json:"apiKey,omitempty"`
		APIKeyID                      string `json:"apiKeyId,omitempty"`
		AccessKey                     string `json:"accessKey,omitempty"`
		SourceIP                      string `json:"sourceIp"`
		CognitoAuthenticationType     string `json:"cognitoAuthenticationType,omitempty"`
		CognitoAuthenticationProvider string `json:"cognitoAuthenticationProvider,omitempty"`
		UserArn                       string `json:"userArn,omitempty"` //nolint: stylecheck
		UserAgent                     string `json:"userAgent"`
		User                          string `json:"user,omitempty"`
	}
	Response struct {
		StatusCode        int                 `json:"statusCode"`
		Headers           map[string]string   `json:"headers"`
		MultiValueHeaders map[string][]string `json:"multiValueHeaders"`
		Body              string              `json:"body"`
		IsBase64Encoded   bool                `json:"isBase64Encoded,omitempty"`
	}
	Handler   func(context.Context, Request) *Response
	BgHandler func(context.Context, Request) error
)

func Start(handler Handler) {
	lambda.Start(func(ctx context.Context, req Request) (*Response, error) {
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
		return Err(fmt.Errorf("server encountered an error while rendering a response: %v", err))
	}
	return TextStatus(status, string(data))
}

func RenderF(payload io.Reader, info any) *Response {
	data, err := io.ReadAll(payload)
	if err != nil {
		return Err(fmt.Errorf("server encountered an error while rendering a response: %v", err))
	}
	return Render(string(data), info)
}

func Render(tmpl string, data any) *Response {
	parsedTmpl := template.Must(template.New("netlify_render").Parse(tmpl))
	var buf bytes.Buffer
	if err := parsedTmpl.Execute(&buf, data); err != nil {
		return Err(fmt.Errorf("server encountered an error while rendering a response: %v", err))
	}
	return Text(buf.String())
}
