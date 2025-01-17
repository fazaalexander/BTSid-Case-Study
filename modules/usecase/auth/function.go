package auth

import (
	"errors"

	pw "github.com/fazaalexander/BTSid-Case-Study.git/helper/password"
	vld "github.com/fazaalexander/BTSid-Case-Study.git/helper/validator"
	"github.com/fazaalexander/BTSid-Case-Study.git/middleware/jwt"
	ue "github.com/fazaalexander/BTSid-Case-Study.git/modules/entity"
)

func (ac *authUseCase) Register(request *ue.RegisterRequest) error {
	if err := vld.Validation(request); err != nil {
		return err
	}

	hashedPassword, err := pw.HashPassword(request.Password)
	if err != nil {
		return err
	}

	request.Password = string(hashedPassword)

	err = ac.authRepo.CreateUser(request)
	if err != nil {
		return err
	}

	return nil
}

func (ac *authUseCase) Login(request *ue.LoginRequest) (interface{}, error) {
	if err := vld.Validation(*request); err != nil {
		return nil, err
	}

	response, password, err := ac.authRepo.Login(request.Email)

	if err != nil {
		return nil, errors.New("wrong email or password")
	}

	if err = pw.VerifyPassword(password, request.Password); err != nil {
		return nil, errors.New("wrong Password")
	}

	token, err := jwt.CreateToken(response.ID, response.Email)
	if err != nil {
		return nil, err
	}

	response.Token = token

	return response, nil
}
