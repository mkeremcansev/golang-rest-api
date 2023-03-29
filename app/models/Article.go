package models

type Article struct {
	ID      uint   `gorm:"primary_key" json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
