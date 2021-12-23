package AuthController

import (
	"app/Exceptions"
	"app/base"
	"app/requests"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type AuthController struct {
	DB *base.DB
}

func New(db *base.DB) *AuthController {
	return &AuthController{db}
}

func (c AuthController) Auth(ctx *gin.Context) {

}

func (c AuthController) Register(ctx *gin.Context) {
	var request requests.Register

	if err := ctx.ShouldBindJSON(&request); err != nil {
		e := Exceptions.New(err.(validator.ValidationErrors))
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": e.FormatToFront()})
		return
	}

	ctx.JSON(http.StatusOK, &request)
}

func (c AuthController) CheckAuth(ctx *gin.Context) {

}
