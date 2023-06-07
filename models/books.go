package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Books struct {
	ID            uuid.UUID `json:"id" gorm:"column:id"`
	Title         string    `json:"title" gorm:"column:title"`
	Author        string    `json:"author" gorm:"column:author"`
	Price         float64   `json:"price" gorm:"column:price"`
	PublishedDate time.Time `json:"published_date" gorm:"column:published_date"`
}
