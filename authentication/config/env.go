package config

import (
	"fmt"
	"os"
)

const (
	awsRegionEnvPropertyName   = "AWS_REGION"
	awsClientIdEnvPropertyName = "AWS_CLIENT_ID"
)

func AwsRegionFromEnv() (string, error) {
	var env = os.Getenv(awsRegionEnvPropertyName)
	if env == "" {
		return "", fmt.Errorf("missing mandatory environment variable %s", awsRegionEnvPropertyName)
	}

	return env, nil
}

func AwsClientIdFromEnv() (string, error) {
	var env = os.Getenv(awsClientIdEnvPropertyName)
	if env == "" {
		return "", fmt.Errorf("missing mandatory environment variable %s", awsClientIdEnvPropertyName)
	}

	return env, nil
}
