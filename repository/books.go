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

	/*
	sections := make([]*model.Section, 0)
	if err := db.Select("id,title").Where("book_id = ?", id).Order("sort asc, id desc").Find(&sections).Error; err != nil {
		return nil, err
	}
	return sections, nil
	*/
	sections := make([]*model.Section, 0)
	db.Raw("SELECT SUM(age) FROM users WHERE role = ?"
	WITH RECURSIVE cte_connect_by AS (
		SELECT 1 AS level, CAST(CONCAT('/', name) AS VARCHAR(4000)) AS connect_by_path, s.* 
		  FROM employees s WHERE id = 1
		UNION ALL
		SELECT level + 1 AS level, CONCAT(connect_by_path, '/', s.name) AS connect_by_path, s.* 
		  FROM cte_connect_by r INNER JOIN employees s ON  r.id = s.mng_id
	 )
	 SELECT id, name, mng_id, level, connect_by_path path
	 FROM cte_connect_by
	 ORDER BY id;
	 , "admin").Scan(&age)

	 return section nil
}

func ReadSection(db *gorm.DB, id string) (*model.Section, error) {

	section := &model.Section{}

	if err := db.Where("id = ?", id).First(&section).Error; err != nil {
		return nil, err
	}
	return section, nil

}

func UpdateSection(db *gorm.DB, section *model.Section) error {
	if err := db.Model(&section).Updates(section).Where("id = ?", section.ID).Error; err != nil {
		return err
	}
	return nil
}
