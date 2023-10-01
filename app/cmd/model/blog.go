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
}

type CreateBlogParams struct {
	UserID  int
	Title   string
	Content string
}

func CreateBlog(params CreateBlogParams) (*Blog, error) {
	blog := Blog{
		UserID:  params.UserID,
		Title:   params.Title,
		Content: params.Content,
	}

	err := db.Create(&blog).Error

	return &blog, err
}
