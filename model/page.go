package models

import "linked-page/db"

type Page struct {
	ID uint `gorm:"primary_key" json:"id,omitempty"`
}

func findPageById(id uint) *Page {
	var page Page
	db.DB.First(&page, id)
	return &page
}
