package models

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	Title    string `gorm:"type:varchar(255);not null"`
	Content  string `gorm:"type:text"`
	Archived bool   `gorm:"default:false"`
}
