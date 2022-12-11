package model

type PageDto struct {
	Title   string `json:"title,omitempty"`
	Content struct {
		Clients     []*ClientDto     `json:"clients,omitempty"`
		Client      *ClientDto       `json:"client,omitempty"`
		Collections []*CollectionDto `json:"collections,omitempty"`
		Collection  *CollectionDto   `json:"collection,omitempty"`
		Books       []*BookDto       `json:"books,omitempty"`
		Book        *BookDto         `json:"book,omitempty"`
	} `json:"content"`
}
