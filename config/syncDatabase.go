package config

import "github.com/orket-sam/go-jwt/models"

func SyncDB() {
	DB.AutoMigrate(&models.User{})
}
