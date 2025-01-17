package auth

import (
	ae "github.com/fazaalexander/BTSid-Case-Study.git/modules/entity"
	"gorm.io/gorm"
)

type AuthRepo interface {
	CreateUser(user *ae.RegisterRequest) error
	Login(email string) (*ae.LoginResponse, string, error)
}

type authRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) AuthRepo {
	return &authRepo{
		db,
	}
}
