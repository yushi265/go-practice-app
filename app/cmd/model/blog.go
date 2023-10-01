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

type CreateBlogParams struct {
	UserID  int
	Title   string
	Content string
}

func CreateBlog(params CreateBlogParams, user User) (*Blog, error) {
	blog := Blog{
		UserID:  params.UserID,
		Title:   params.Title,
		Content: params.Content,
		User:    user,
	}

	err := db.Preload("User").Create(&blog).Error

	return &blog, err
}

func GetBlogs(userID string) ([]Blog, error) {
	var blogs []Blog

	var err error

	if userID != "" {
		err = db.Preload("User").Where("user_id = ?", userID).Find(&blogs).Error
	} else {
		err = db.Preload("User").Find(&blogs).Error
	}

	return blogs, err
}
