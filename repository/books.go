package repository

import (
	"blubooks/model"

	"gorm.io/gorm"
)

func ListBooks(db *gorm.DB, id string) (model.Books, error) {

	books := make([]*model.Book, 0)
	if err := db.Where("collection_id = ?", id).Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func ReadBook(db *gorm.DB, id string) (*model.Book, error) {

	book := &model.Book{}

	if err := db.Where("id = ?", id).First(&book).Error; err != nil {
		return nil, err
	}
	return book, nil

}

func ListSections(db *gorm.DB, id string) (model.Sections, error) {

	sections := make([]*model.Section, 0)
	if err := db.Where("book_id = ?", id).Order("sort asc, id desc").Find(&sections).Error; err != nil {
		return nil, err
	}
	return sections, nil
}
func ListSectionsTitles(db *gorm.DB, id string) (model.Sections, error) {

	sections := make([]*model.Section, 0)
	if err := db.Select("id,title").Where("book_id = ?", id).Order("sort asc, id desc").Find(&sections).Error; err != nil {
		return nil, err
	}
	return sections, nil
}

func ReadSection(db *gorm.DB, id string) (*model.Section, error) {

	section := &model.Section{}

	if err := db.Where("id = ?", id).First(&section).Error; err != nil {
		return nil, err
	}
	return section, nil

}
