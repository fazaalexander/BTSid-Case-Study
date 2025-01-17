package auth

import (
	"errors"
	"time"

	ue "github.com/fazaalexander/BTSid-Case-Study.git/modules/entity"
)

func (ar *authRepo) CreateUser(user *ue.RegisterRequest) error {
	existingUser := ue.User{}
	if err := ar.db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return errors.New("email already exists")
	}

	userTable := ue.User{
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: time.Now(),
	}

	if err := ar.db.Create(&userTable).Error; err != nil {
		return err
	}

	return nil
}

func (ar *authRepo) Login(email string) (*ue.LoginResponse, string, error) {
	user := &ue.User{}
	if err := ar.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, "", errors.New("record not found")
	}

	response := &ue.LoginResponse{
		ID:    user.ID,
		Email: user.Email,
	}

	return response, user.Password, nil
}
