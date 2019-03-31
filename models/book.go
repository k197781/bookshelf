package models

import (
	"time"
)

type Book struct {
	Id int64 `gorm:"primary_key" json:"id"`
	Title string `sql:"size:255" json:"title"`
	Author string `sql:"size:255" json:"author"`
	Publisher string `sql:"size:255" json:"publisher"`
	PublishedDate time.Time `json:"published_date"`
	CreatedAt time.Time
	UpdatedAT time.Time
	DeletedAt *time.Time `sql:"index" json:"-"`
}
