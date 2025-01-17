package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	*gorm.Model `json:"-"`
	ID          uint64    `gorm:"primarykey" json:"UserId"`
	Username    string    `validate:"required"`
	Email       string    `validate:"required"`
	Password    string    `validate:"required,min=8"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type RegisterRequest struct {
	Username string `json:"Username" form:"Username" validate:"required"`
	Email    string `json:"Email" form:"Email" validate:"required,email"`
	Password string `json:"Password" form:"Password" validate:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"Email" form:"Email" validate:"required,email"`
	Password string `json:"Password" form:"Password" validate:"required,min=8"`
}

type LoginResponse struct {
	ID    uint64 `json:"Id" form:"Id"`
	Email string `json:"Email" form:"Email"`
	Token string `json:"Token" form:"Token"`
}
