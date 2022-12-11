package repository

import (
	"blubooks/model"

	"gorm.io/gorm"
)

func ListCollectionByClientID(db *gorm.DB, id string) (model.Collections, error) {

	collections := make([]*model.Collection, 0)
	if err := db.Where("client_id = ?", id).Find(&collections).Error; err != nil {
		return nil, err
	}
	return collections, nil
}

func ListBooks(db *gorm.DB, id string) (model.Books, error) {

	books := make([]*model.Book, 0)
	if err := db.Where("collection_id = ?", id).Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func ReadCollection(db *gorm.DB, id string) (*model.Collection, error) {

	collection := &model.Collection{}

	if err := db.Where("id = ?", id).First(&collection).Error; err != nil {
		return nil, err
	}
	return collection, nil

}
