package handlers

import (
	"context"
	"errors"

	"github.com/sithsithsith/mugi-cognito/internal/services"
	"github.com/sithsithsith/mugi-cognito/internal/utils"
)

type SignUpHandler struct {
	CognitoService *services.CognitoService
}

func NewSignUpHandler(cs *services.CognitoService) *SignUpHandler {
	return &SignUpHandler{CognitoService: cs}
}

func (sh *SignUpHandler) LambdaHandler(ctx context.Context, event map[string]interface{}) (map[string]interface{}, error) {
	phoneNumber, ok1 := event["phone_number"].(string)
	password, ok2 := event["password"].(string)
	if !ok1 || !ok2 {
		return utils.LambdaErrorResponse(errors.New("missing required fields")), nil
	}

	err := sh.CognitoService.SignUp(phoneNumber, password)
	if err != nil {
		return utils.LambdaErrorResponse(err), nil
	}

	return utils.LambdaSuccessResponse("Sign-up successful"), nil
}
