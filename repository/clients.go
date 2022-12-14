package repository

import (
	"blubooks/model"

	"gorm.io/gorm"
)

func ListClients(db *gorm.DB) (model.Clients, error) {

	clients := make([]*model.Client, 0)
	if err := db.Find(&clients).Error; err != nil {
		return nil, err
	}
	return clients, nil
}

func ReadClient(db *gorm.DB, id string) (*model.Client, error) {

	client := &model.Client{}

	if err := db.Where("id = ?", id).First(&client).Error; err != nil {
		return nil, err
	}
	return client, nil

}
