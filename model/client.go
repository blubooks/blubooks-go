package model

import (
	g "blubooks/adapter/gorm"
)

type Clients []*Client
type Client struct {
	g.ModelUID
	Title string
}

type ClientDtos []*ClientDto
type ClientDto struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func (o Client) ToDto() *ClientDto {
	return &ClientDto{
		ID:    o.ID,
		Title: o.Title,
	}
}
func (os Clients) ToDto() ClientDtos {
	dtos := make([]*ClientDto, len(os))
	for i, o := range os {
		dtos[i] = o.ToDto()
	}
	return dtos
}
