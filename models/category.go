package models

import "github.com/gofrs/uuid"

type Category struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Name string    `gorm:"type:varchar(255);not null" json:"name"`
}
