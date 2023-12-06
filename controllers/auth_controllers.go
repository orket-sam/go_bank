package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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

	// Retrive user
	var user models.User
	config.DB.Find(&user, "email=?", body.Email)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	// compare
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Invalid password",
		})
		return
	}

	// generate token

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email, "exp": time.Now().Add(time.Hour * 24).Unix()})

	// Sign and get the complete encoded token as a string using the secret
	key := []byte("Wy0Siv5XGwgewqeubsdaihdhebQE/HDWIOHECBCJD")
	tokenString, err := token.SignedString(key)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	c.JSON(200, gin.H{"token": tokenString})

}
