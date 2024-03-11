package lambda

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
)

type Response struct {
	Data    any    `json:"data,omitempty"`
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Status  int    `json:"status"`
}

func buildErrorResponseBody(status int, message string) []byte {
	responseBody := Response{
		Status:  status,
		Message: message,
		Success: false,
	}
	body, _ := json.Marshal(responseBody)

	return body
}

func buildSuccessResponseBody(data any) []byte {
	responseBody := Response{
		Status:  http.StatusOK,
		Data:    data,
		Success: true,
	}

	body, _ := json.Marshal(responseBody)

	return body
}

func SendError(statusCode int, errorMessage string) (events.APIGatewayProxyResponse, error) {
	responseBody := buildErrorResponseBody(statusCode, errorMessage)

	return events.APIGatewayProxyResponse{
		Headers:    map[string]string{"Content-Type": "application/json"},
		StatusCode: statusCode,
		Body:       string(responseBody),
	}, nil
}

func SendValidationError(statusCode int, validationMessage string) (events.APIGatewayProxyResponse, error) {
	responseBody := buildErrorResponseBody(statusCode, validationMessage)

	return events.APIGatewayProxyResponse{
		Headers:    map[string]string{"Content-Type": "application/json"},
		StatusCode: statusCode,
		Body:       string(responseBody),
	}, nil
}

func Send(statusCode int, data any) (events.APIGatewayProxyResponse, error) {

	responseBody := buildSuccessResponseBody(data)

	return events.APIGatewayProxyResponse{
		Headers:    map[string]string{"Content-Type": "application/json"},
		StatusCode: statusCode,
		Body:       string(responseBody),
	}, nil
}
