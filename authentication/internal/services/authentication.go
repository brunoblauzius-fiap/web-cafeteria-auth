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
	RespondToAuthChallenge(ctx context.Context, params *cognitoidentityprovider.RespondToAuthChallengeInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.RespondToAuthChallengeOutput, error)
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

func (s Authentication) Auth(ctx context.Context, request models.AuthenticationRequest) (*models.AuthenticationResponse, error) {
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
		return nil, err
	}

	authenticationResult := authResult.AuthenticationResult

	if authenticationResult != nil {
		response := models.AuthenticationResponse{
			AccessToken: *authenticationResult.AccessToken,
			ExpiresIn:   *&authenticationResult.ExpiresIn,
			TokenType:   *authenticationResult.TokenType,
		}

		return &response, nil
	}

	if authResult.ChallengeName == "NEW_PASSWORD_REQUIRED" {
		respondInput := &cognitoidentityprovider.RespondToAuthChallengeInput{
			ChallengeName: types.ChallengeNameTypeNewPasswordRequired,
			ChallengeResponses: map[string]string{
				"NEW_PASSWORD": request.Password,
				"USERNAME":     request.Email,
			},
			ClientId: &clientId,
			Session:  &*authResult.Session,
		}
		authResult, err := s.cognitoClient.RespondToAuthChallenge(ctx, respondInput)
		if err != nil {
			return nil, err
		}

		authenticationResult := authResult.AuthenticationResult
		response := models.AuthenticationResponse{
			AccessToken: *authenticationResult.AccessToken,
			ExpiresIn:   *&authenticationResult.ExpiresIn,
			TokenType:   *authenticationResult.TokenType,
		}

		return &response, nil
	}

	return nil, nil
}
