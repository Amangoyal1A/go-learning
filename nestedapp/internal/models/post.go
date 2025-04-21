package models

type Post struct {
	ID     uint     `gorm:"primaryKey"`
	Title  string   `json:"title"`
	Body   string   `json:"body"`
	UserID uint
	Comments []Comment `gorm:"constraint:OnDelete:CASCADE"`
}
