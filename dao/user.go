package dao

import (
	"errors"
	"temphtmlmounter/model"

	"gorm.io/gorm"
)

func CreateUser(user *model.User) error {
	return db.Create(user).Error
}

func GetUserByID(id uint) (*model.User, error) {
	var user model.User
	err := db.First(&user, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

func GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	err := db.Where("user_name = ?", username).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}