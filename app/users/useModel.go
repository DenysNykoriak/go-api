package app_users

import "gorm.io/gorm"

type User struct {
	gorm.Model

	FirstName      string `json:"first_name" gorm:"not null"`
	LastName       string `json:"last_name" gorm:"not null"`
	Email          string `json:"email" gorm:"unique;not null"`
	HashedPassword string `json:"hashed_password" gorm:"not null"`
}
