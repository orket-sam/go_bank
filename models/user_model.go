package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `json:"user_name" gorm:"unique" binding:"required"`
	Email    string `json:"email"  gorm:"unique" binding:"email"`
	Password string `json:"password" binding:"required"`
}
