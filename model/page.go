package model

const (
	TYPE_Clients     string = "clients"
	TYPE_Client             = "client"
	TYPE_Collections        = "collections"
	TYPE_Collection         = "collection"
	TYPE_Books              = "books"
	TYPE_Book               = "book"
)

type PageBackLink struct {
	Type string `json:"type,omitempty"`
}
type PageDto struct {
	Title   string `json:"title,omitempty"`
	Type    string `json:"type,omitempty"`
	Back    string `json:"back,omitempty"`
	Content struct {
		Clients     []*ClientDto     `json:"clients,omitempty"`
		Client      *ClientDto       `json:"client,omitempty"`
		Collections []*CollectionDto `json:"collections,omitempty"`
		Collection  *CollectionDto   `json:"collection,omitempty"`
		Books       []*BookDto       `json:"books,omitempty"`
		Book        *BookDto         `json:"book,omitempty"`
	} `json:"content"`
}
