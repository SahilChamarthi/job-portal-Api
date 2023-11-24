package repository

import (
	"errors"
	"project/internal/model"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (r *Repo) CreateUser(u model.User) (model.User, error) {
	err := r.db.Create(&u).Error
	if err != nil {
		return model.User{}, err
	}
	return u, nil
}

func (r *Repo) FetchUserByEmail(s string) (model.User, error) {
	var u model.User
	tx := r.db.Where("email=?", s).First(&u)
	if tx.Error != nil {
		return model.User{}, nil
	}
	return u, nil

}

func (r *Repo) FindEmail(s string) (bool, error) {
	var u model.User
	tx := r.db.Where("email=?", s).First(&u)
	if tx.Error != nil {
		return false, tx.Error
	}
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return false, tx.Error
	}
	return true, nil
}

func (r *Repo) UpdatePassword(email string, newPass string) (bool, error) {

	var u model.User
	tx := r.db.Where("email=?", email).First(&u)

	if tx.Error != nil {
		return false, tx.Error
	}
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return false, tx.Error
	}
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(newPass), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Msg("error occured in hashing password")
		return false, errors.New("hashing password failed")
	}
	u.PasswordHash = string(hashedPass)
	if err := r.db.Save(&u).Error; err != nil {
		return false, err
	}

	return true, nil
}
