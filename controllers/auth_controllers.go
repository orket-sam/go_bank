package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/orket-sam/go-jwt/config"
	"github.com/orket-sam/go-jwt/models"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return

	}

	// Hash password

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Failed to hash password"})
		return

	}

	user.Password = string(hashedPassword)

	result := config.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": (result.Error.Error())})

		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "User created succesfuly"})

}

func SignIn(c *gin.Context) {
	var body struct {
		Email    string `json:"email" binding:"email"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

}
