package models

type Comment struct {
	ID     uint   `gorm:"primaryKey"`
	Text   string `json:"text"`
	PostID uint
}