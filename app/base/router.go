package base

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func LoadRouter() Router {
	var r = Router{
		Engine: gin.Default(),
	}

	r.Static("storage/screenshot/", "./storage/screenshot/")

	return r
}
