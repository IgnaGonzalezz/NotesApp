package models

import "time"

type Note struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`

	Title      string
	Content    string
	Archived   bool
	Categories []Category `gorm:"many2many:note_categories;"`
}
