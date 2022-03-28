package requests

import (
	"app/exceptions/ValidationExeption"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseRequest struct {
}

func (r *BaseRequest) Validate(ctx *gin.Context) bool {
	return Validate(r, ctx)
}

func Validate(r interface{}, ctx *gin.Context) bool {
	err := ctx.ShouldBind(r)

	if err == nil {
		return true
	}

	e := ValidationExeption.New(err)
	ctx.JSON(http.StatusBadRequest, gin.H{"errors": e.FormatToFront()})
	return false
}
