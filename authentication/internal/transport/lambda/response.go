package lambda

import (
	"encoding/json"

	"github.com/4SOAT/web-cafeteria-auth/authentication/internal/models"
	"github.com/aws/aws-lambda-go/events"
)

type Response struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int32  `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func buildErrorResponseBody(message string) []byte {
	responseBody := ErrorResponse{
		Message: message,
	}
	body, _ := json.Marshal(responseBody)

	return body
}

func buildSuccessResponseBody(response models.AuthenticationResponse) []byte {
	responseBody := Response{
		AccessToken: response.AccessToken,
		ExpiresIn:   response.ExpiresIn,
		TokenType:   response.TokenType,
	}

	body, _ := json.Marshal(responseBody)

	return body
}

func SendError(statusCode int, errorMessage string) (events.APIGatewayProxyResponse, error) {
	responseBody := buildErrorResponseBody(errorMessage)

	return events.APIGatewayProxyResponse{
		Headers:    map[string]string{"Content-Type": "application/json"},
		StatusCode: statusCode,
		Body:       string(responseBody),
	}, nil
}

func SendValidationError(statusCode int, validationMessage string) (events.APIGatewayProxyResponse, error) {
	responseBody := buildErrorResponseBody(validationMessage)

	return events.APIGatewayProxyResponse{
		Headers:    map[string]string{"Content-Type": "application/json"},
		StatusCode: statusCode,
		Body:       string(responseBody),
	}, nil
}

func Send(statusCode int, response models.AuthenticationResponse) (events.APIGatewayProxyResponse, error) {
	responseBody := buildSuccessResponseBody(response)

	return events.APIGatewayProxyResponse{
		Headers:    map[string]string{"Content-Type": "application/json"},
		StatusCode: statusCode,
		Body:       string(responseBody),
	}, nil
}
