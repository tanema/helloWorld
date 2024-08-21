package netlify

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
		CognitoIdentityPoolID         string `json:"cognitoIdentityPoolId,omitempty"`
		AccountID                     string `json:"accountId,omitempty"`
		CognitoIdentityID             string `json:"cognitoIdentityId,omitempty"`
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
)
