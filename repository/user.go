package repository

import (
	"blubooks/model"

	"gorm.io/gorm"
)

func ReadUser(db *gorm.DB, id string) (*model.User, error) {
	user := &model.User{}

	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByEmail(db *gorm.DB, email string) (*model.User, error) {
	user := &model.User{}

	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
