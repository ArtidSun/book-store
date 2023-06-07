package models

import "github.com/gofrs/uuid"

type BookCategory struct {
	BookID     uuid.UUID `json:"book_id gorm:"column:book_id"`
	CategoryID uuid.UUID `json:"category_id" gorm:"column:category_id"`

	Book     Books    `gorm:"foreignKey:BookID;references:ID"`
	Category Category `gorm:"foreignKey:CategoryID;references:ID"`
}
