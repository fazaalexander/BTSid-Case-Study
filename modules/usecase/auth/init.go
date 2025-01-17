package auth

import (
	ae "github.com/fazaalexander/BTSid-Case-Study.git/modules/entity"
	ar "github.com/fazaalexander/BTSid-Case-Study.git/modules/repository/auth"
)

type AuthUseCase interface {
	Login(request *ae.LoginRequest) (interface{}, error)
	Register(user *ae.RegisterRequest) error
}

type authUseCase struct {
	authRepo ar.AuthRepo
}

func New(authRepo ar.AuthRepo) *authUseCase {
	return &authUseCase{
		authRepo,
	}
}
