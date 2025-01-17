package auth

import (
	ac "github.com/fazaalexander/BTSid-Case-Study.git/modules/usecase/auth"
)

type AuthHandler struct {
	authUseCase ac.AuthUseCase
}

func New(authUseCase ac.AuthUseCase) *AuthHandler {
	return &AuthHandler{
		authUseCase,
	}
}
