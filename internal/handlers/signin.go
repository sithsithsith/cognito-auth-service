package handlers

import (
	"context"
	"errors"

	"github.com/sithsithsith/mugi-cognito/internal/services"
	"github.com/sithsithsith/mugi-cognito/internal/utils"
)

type SignInHandler struct {
	CognitoService *services.CognitoService
}

func NewSignInHandler(cs *services.CognitoService) *SignInHandler {
	return &SignInHandler{CognitoService: cs}
}

func (sh *SignInHandler) LambdaHandler(ctx context.Context, event map[string]interface{}) (map[string]interface{}, error) {
	phoneNumber, ok1 := event["phone_number"].(string)
	password, ok2 := event["password"].(string)
	if !ok1 || !ok2 {
		return utils.LambdaErrorResponse(errors.New("missing required fields")), nil
	}

	token, err := sh.CognitoService.SignIn(phoneNumber, password)
	if err != nil {
		return utils.LambdaErrorResponse(err), nil
	}

	return utils.LambdaSuccessResponse(map[string]string{"token": token}), nil
}
