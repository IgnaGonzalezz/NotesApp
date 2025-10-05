package models

import "time"

type Category struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`

	Name  string `gorm:"unique;not null"`
	Notes []Note `gorm:"many2many:note_categories;"`
}
