package main

import (
	"github.com/sithsithsith/mugi-cognito/internal/config"
	"github.com/sithsithsith/mugi-cognito/internal/handlers"
	"github.com/sithsithsith/mugi-cognito/internal/services"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	client := config.GetCognitoClient()
	cognitoService := services.NewCognitoService(client)
	signUpHandler := handlers.NewSignUpHandler(cognitoService)

	// Start Lambda handler
	lambda.Start(signUpHandler.LambdaHandler)
}
