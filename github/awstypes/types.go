package awstypes

import "github.com/aws/aws-lambda-go/events"

// APIGatewayProxyResponse configures the response to be returned by API Gateway for the request
type APIGatewayProxyResponse struct {
	StatusCode      int               `json:"statusCode"`
	Headers         map[string]string `json:"headers"`
	Body            string            `json:"body"`
	IsBase64Encoded bool              `json:"isBase64Encoded,omitempty"`
}

type Response APIGatewayProxyResponse
type Request events.APIGatewayProxyRequest
