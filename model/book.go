package model

import (
	g "blubooks/adapter/gorm"
	"bytes"

	"github.com/yuin/goldmark"
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

type Sections []*Section
type Section struct {
	g.ModelUID
	Title     string
	Content   string
	BookID    string
	SectionID string
}

type SectionDtos []*SectionDto
type SectionDto struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content,omitempty"`
}

func (o Section) ToMarkdownDto() *SectionDto {

	var buffer bytes.Buffer
	var content = ""
	if err := goldmark.Convert([]byte(o.Content), &buffer); err != nil {
		content = "Fehler beim Parsen"
	} else {
		content = buffer.String()
	}

	return &SectionDto{
		ID:      o.ID,
		Title:   o.Title,
		Content: content,
	}
}
func (o Section) ToDto() *SectionDto {
	return &SectionDto{
		ID:      o.ID,
		Title:   o.Title,
		Content: o.Content,
	}
}

func (os Sections) ToDto() SectionDtos {
	dtos := make([]*SectionDto, len(os))
	for i, o := range os {
		dtos[i] = o.ToDto()
	}
	return dtos
}
