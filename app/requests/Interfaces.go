package requests

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/guregu/null.v4"
)

type ValidateRequest interface {
	Validate(ctx *gin.Context) bool
}

type DialRequestInterface interface {
	GetUrl() string
	GetDescription() null.String
	GetName() null.String
}
