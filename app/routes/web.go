package routes

import (
	"app/controllers"
	"github.com/gin-gonic/gin"
)

func (r Route) Web() {
	authController := controllers.NewAuth(r.DB)

	r.Router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	auth := r.Router.Group("auth")
	{
		auth.POST("register", authController.Register)
		auth.POST("/", authController.Auth)
	}
}
