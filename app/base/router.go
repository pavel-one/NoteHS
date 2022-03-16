package base

import (
	"app/validations"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

type Router struct {
	*gin.Engine
}

func LoadRouter() Router {
	gin.SetMode(os.Getenv("GIN_MODE"))

	var r = Router{
		Engine: gin.Default(),
	}

	validations.SetNullValidators()

	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}))

	return r
}
