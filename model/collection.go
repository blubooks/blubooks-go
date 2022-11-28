package model

import (
	g "blubooks/adapter/gorm"
)

type Collections []*Collection
type Collection struct {
	g.ModelUID
	Title    string
	ClientID string
	Client   Client
}

type CollectionDtos []*CollectionDto
type CollectionDto struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func (o Collection) ToDto() *CollectionDto {
	return &CollectionDto{
		ID:    o.ID,
		Title: o.Title,
	}
}
func (os Collections) ToDto() CollectionDtos {
	dtos := make([]*CollectionDto, len(os))
	for i, o := range os {
		dtos[i] = o.ToDto()
	}
	return dtos
}
