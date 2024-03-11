package main

import (
	"context"
	"fmt"
	"log"

	internalConfig "github.com/4SOAT/web-cafeteria-auth/authentication/config"
	handler "github.com/4SOAT/web-cafeteria-auth/authentication/internal/handlers/lambda"
	"github.com/4SOAT/web-cafeteria-auth/authentication/internal/logging"
	"github.com/4SOAT/web-cafeteria-auth/authentication/internal/models"
	"github.com/4SOAT/web-cafeteria-auth/authentication/internal/services"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/joho/godotenv"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

// used for testing
func main2() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()
	logger, err := logging.New()
	if err != nil {
		logger.Fatal(err.Error())
	}

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-east-1"))
	if err != nil {
		logger.Fatal(err.Error())
	}

	cognitoClient := cognitoidentityprovider.NewFromConfig(cfg)

	srv := services.New(logger, cognitoClient)

	token, err := srv.Auth(ctx, models.AuthenticationRequest{
		Email:    "primeirinho@gmail.com",
		Password: "Primeiro@123",
	})

	if err != nil {
		logger.Fatal(err.Error())
	}

	fmt.Println(token)
}

func run() error {

	ctx := context.Background()
	logger, err := logging.New()
	if err != nil {
		return fmt.Errorf("error loading log config: %s", err)
	}

	awsRegion, err := internalConfig.AwsRegionFromEnv()
	if err != nil {
		logger.Fatal(err.Error())
	}

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(awsRegion))
	if err != nil {
		logger.Fatal(err.Error())
	}

	cognitoClient := cognitoidentityprovider.NewFromConfig(cfg)
	srv := services.New(logger, cognitoClient)

	lambda.Start(handler.New(logger, srv).Handler)

	return nil
}
