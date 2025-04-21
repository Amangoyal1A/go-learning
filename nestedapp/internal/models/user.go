package models

type User struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Posts []Post `gorm:"constraint:OnDelete:CASCADE" json:"posts"`
}
