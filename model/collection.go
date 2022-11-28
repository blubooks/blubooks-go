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
