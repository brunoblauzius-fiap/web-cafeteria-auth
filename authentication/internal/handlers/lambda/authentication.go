package lambda

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/4SOAT/web-cafeteria-auth/authentication/internal/logging"
	"github.com/4SOAT/web-cafeteria-auth/authentication/internal/models"
	transport "github.com/4SOAT/web-cafeteria-auth/authentication/internal/transport/lambda"
	"github.com/aws/aws-lambda-go/events"
)

const (
	ErrorInvalidEmail       = "Invalid email"
	ErrorInvalidRequestBody = "Failed decode request body"
)

type Service interface {
	Auth(ctx context.Context, request models.AuthenticationRequest) (string, error)
}

type Authentication struct {
	logger  logging.Logger
	service Service
}

func New(log logging.Logger, srv Service) Authentication {
	return Authentication{
		logger:  log,
		service: srv,
	}
}

func (s Authentication) Handler(e events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var request models.AuthenticationRequest

	if err := json.Unmarshal([]byte(e.Body), &request); err != nil {
		return transport.SendError(http.StatusInternalServerError, ErrorInvalidRequestBody)
	}

	s.logger.Info("Handling Authentication event")

	if request.Email == "" {
		return transport.SendValidationError(http.StatusBadRequest, ErrorInvalidEmail)
	}

	token, err := s.service.Auth(context.Background(), request)
	if err != nil {
		return transport.SendError(http.StatusInternalServerError, err.Error())
	}

	return transport.Send(http.StatusOK, token)
}
