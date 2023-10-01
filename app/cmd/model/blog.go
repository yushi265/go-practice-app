package model

import (
	"time"

	"gorm.io/gorm"
)

type Blog struct {
	ID        int `gorm:"primaryKey;type:BIGINT UNSIGNED"`
	UserID    int
	Title     string `gorm:"size:255"`
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	User      User
}
