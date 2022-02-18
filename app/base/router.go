package base

import (
	"app/validations"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func LoadRouter() Router {
	var r = Router{
		Engine: gin.Default(),
	}

	r.Use(cors.Default())
	validations.SetNullValidators()

	r.Static("storage/screenshot/", "./storage/screenshot/")

	return r
}
