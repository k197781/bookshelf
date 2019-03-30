package models

import (
	"time"
)

type Book struct {
	Id int64
	Title string `sql:"size:255"`
	Author string `sql:"size:255"`
	Publisher string `sql:"size:255"`
	PublishedDate string `sql:"size:255"`
	CreatedAt time.Time
	UpdatedAT time.Time
	DeletedAt time.Time
}
