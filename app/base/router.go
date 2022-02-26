package base

import (
	"app/validations"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

type Router struct {
	*gin.Engine
}

func LoadRouter() Router {
	var r = Router{
		Engine: gin.Default(),
	}

	validations.SetNullValidators()

	r.Static("storage/screenshot/", "./storage/screenshot/")

	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}))

	return r
}
