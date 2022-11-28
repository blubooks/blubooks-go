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
