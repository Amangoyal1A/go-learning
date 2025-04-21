package models

type Post struct {
	ID       uint      `gorm:"primaryKey" json:"id"`
	Title    string    `json:"title" binding:"required"`
	UserID   uint      `json:"user_id" binding:"required"`
	User     *User      `gorm:"constraint:OnDelete:CASCADE" json:"-"`
	Comments []Comment `gorm:"constraint:OnDelete:CASCADE" json:"comments"`
}
