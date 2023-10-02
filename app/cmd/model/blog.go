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

type UpdateBlogParams struct {
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

func PutBlog(blogID int, params UpdateBlogParams) (*Blog, error) {
	var blog Blog

	err := db.Model(&Blog{}).Where("id = ?", blogID).Update("title", params.Title).Update("content", params.Content).Error

	if err != nil {
		return &blog, err
	}

	err = db.Preload("User").Where("id = ?", blogID).First(&blog).Error

	return &blog, err
}

func DeleteBlog(blogID int) error {
	return db.Delete(&Blog{}, blogID).Error
}

func GetBlog(blogID int) (*Blog, error) {
	var blog *Blog

	err := db.Preload("User").Where("id = ?", blogID).First(&blog).Error

	return blog, err
}
