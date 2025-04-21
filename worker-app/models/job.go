package models

import (
	"time"

	"gorm.io/gorm"
)

type Job struct {
	ID        uint           `gorm:"primaryKey"`
	Task      string         `json:"task"`
	Status    string         `json:"status"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
