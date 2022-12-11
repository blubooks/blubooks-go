package model

import (
	g "blubooks/adapter/gorm"
)

type Books []*Book
type Book struct {
	g.ModelUID
	Title        string
	CollectionID string
	Collection   Collection
}

type BookDtos []*BookDto
type BookDto struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func (o Book) ToDto() *BookDto {
	return &BookDto{
		ID:    o.ID,
		Title: o.Title,
	}
}
func (os Books) ToDto() BookDtos {
	dtos := make([]*BookDto, len(os))
	for i, o := range os {
		dtos[i] = o.ToDto()
	}
	return dtos
}
