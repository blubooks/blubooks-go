package model

import (
	g "blubooks/adapter/gorm"
	"bytes"
	"strings"

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
type SectionNavis []*SectionNavi
type SectionNavi struct {
	g.ModelUID
	Title     string
	Content   string
	BookID    string
	SectionID string
	Level     int
	IDs       string
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

type SectionNaviDtos []*SectionNaviDto
type SectionNaviDto struct {
	ID        string          `json:"id"`
	Title     string          `json:"title"`
	Level     int             `json:"level,omitempty"`
	SectionID string          `json:"section_id,omitempty"`
	IDs       []string        `json:"ids,omitempty"`
	Children  SectionNaviDtos `json:"children,omitempty"`
}

func (o SectionNavi) ToSectionNaviDto() *SectionNaviDto {
	return &SectionNaviDto{
		ID:        o.ID,
		SectionID: o.SectionID,
		Title:     o.Title,
		Level:     o.Level,
		IDs:       strings.Split(o.IDs, ","),
	}
}
func (os SectionNavis) ToSectionNaviDto() SectionNaviDtos {

	var result SectionNaviDtos

	for _, val := range os {
		res := val.ToSectionNaviDto()
		var found bool

		if val.SectionID != "" {

			for _, root := range result {
				parent := findById(root, val.SectionID)
				if parent != nil {
					parent.Children = append(parent.Children, res)
					found = true
					break
				}
			}
		}

		if !found {
			result = append(result, res)
		}

	}

	return result
}

func findById(root *SectionNaviDto, id string) *SectionNaviDto {
	queue := make([]*SectionNaviDto, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		nextUp := queue[0]
		queue = queue[1:]
		if nextUp.ID == id {
			return nextUp
		}
		if len(nextUp.Children) > 0 {
			for _, child := range nextUp.Children {
				queue = append(queue, child)
			}
		}
	}
	return nil
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
