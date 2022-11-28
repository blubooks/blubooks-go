package repository

import (
	"blubooks/model"

	"gorm.io/gorm"
)

func ListCollection(db *gorm.DB, id string) (model.Collections, error) {

	collections := make([]*model.Collection, 0)
	if err := db.Where("client_id = ?", id).Find(&collections).Error; err != nil {
		return nil, err
	}
	return collections, nil
}
