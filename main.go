package main

import (
	"github.com/gin-gonic/gin"
	"github.com/orket-sam/go-jwt/config"
	"github.com/orket-sam/go-jwt/routes"
)

func init() {
	config.ConnectDb()
	config.SyncDB()
}

func main() {
	r := gin.Default()
	routes.AuthRoutes(r)
	r.Run("localhost:8080")
}
