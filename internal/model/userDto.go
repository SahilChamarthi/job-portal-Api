package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName     string `json:"name" validate:"required,unique" gorm:"unique,notnull"`
	Email        string `json:"email" validate:"required"`
	PasswordHash string `json:"-" validate:"required"`
}

type UserSignup struct {
	UserName string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type PasswordReset struct {
	Otp             string `json:"otp" validate:"required"`
	Email           string `json:"email" validate:"required"`
	NewPassword     string `json:"new_password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}
