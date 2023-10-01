package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int    `gorm:"primaryKey;type:BIGINT UNSIGNED"`
	Name      string `gorm:"size:255"`
	Email     string `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func GetUsers() ([]User, error) {
	var users []User

	err := db.Find(&users).Error

	return users, err
}

func CreateUser(name, email string) (*User, error) {
	user := User{
		Name:  name,
		Email: email,
	}

	err := db.Create(&user).Error

	return &user, err
}

func GetUser(userID int) (*User, error) {
	var user User

	err := db.Where("id = ?", userID).First(&user).Error

	return &user, err
}

func UpdateUser(userID int, name string) (*User, error) {
	var user User

	err := db.Model(&User{}).Where("id = ?", userID).Update("name", name).Error

	if err != nil {
		return &user, err
	}

	err = db.Where("id = ?", userID).First(&user).Error

	return &user, err
}

func DeleteUser(userID int) error {
	err := db.Delete(&User{}, userID).Error

	return err
}
