package models

type Comment struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Body   string `json:"body" binding:"required"`
	PostID uint   `json:"post_id" binding:"required"`
	UserID uint   `json:"user_id" binding:"required"`
	Post   *Post   `gorm:"constraint:OnDelete:CASCADE" json:"-"`
	User   *User   `gorm:"constraint:OnDelete:CASCADE" json:"-"`
}
