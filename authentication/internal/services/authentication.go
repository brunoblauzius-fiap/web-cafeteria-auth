package services

import (
	"context"

	internalConfig "github.com/4SOAT/web-cafeteria-auth/authentication/config"
	"github.com/4SOAT/web-cafeteria-auth/authentication/internal/logging"
	"github.com/4SOAT/web-cafeteria-auth/authentication/internal/models"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

type CognitoClient interface {
	InitiateAuth(ctx context.Context, params *cognitoidentityprovider.InitiateAuthInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.InitiateAuthOutput, error)
}

type Authentication struct {
	logger        logging.Logger
	cognitoClient CognitoClient
}

func New(log logging.Logger, c CognitoClient) Authentication {
	return Authentication{
		logger:        log,
		cognitoClient: c,
	}
}

func (s Authentication) Auth(ctx context.Context, request models.AuthenticationRequest) (string, error) {
	s.logger.Info("Serving Authentication event", logging.String("email", request.Email))
	clientId, err := internalConfig.AwsClientIdFromEnv()
	if err != nil {
		s.logger.Fatal(err.Error())
	}

	authInput := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: types.AuthFlowTypeUserPasswordAuth,
		AuthParameters: map[string]string{
			"USERNAME": request.Email,
			"PASSWORD": request.Password,
		},
		ClientId: &clientId,
	}

	authResult, err := s.cognitoClient.InitiateAuth(ctx, authInput)
	if err != nil {
		return "", err
	}

	token := *authResult.Session

	return token, nil
}
