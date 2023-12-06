package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/orket-sam/go-jwt/controllers"
)

func AuthRoutes(r *gin.Engine) {
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Welcome to gin")
	})

	r.POST("signUp", controllers.SignUp)
	r.POST("/signIn", controllers.SignIn)
}
