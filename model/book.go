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
type SectionForm struct {
	Title   string `json:"title" form:"required,max=255"`
	Content string `json:"content"`
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

func (f *SectionForm) ToModel() (*Section, error) {
	/*pubDate, err := time.Parse("2006-01-02", f.PublishedDate)
	if err != nil {
		return nil, err
	}
	*/
	return &Section{
		Title:   f.Title,
		Content: f.Content,
	}, nil
}
