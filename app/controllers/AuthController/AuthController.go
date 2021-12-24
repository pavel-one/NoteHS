package AuthController

import (
	"app/base"
	"app/exceptions/ValidationExeption"
	"app/models"
	"app/requests"
	"github.com/gin-gonic/gin"
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
		e := ValidationExeption.New(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": e.FormatToFront()})
		return
	}

	var u = models.User{
		Username: request.Email,
		Email:    request.Email,
		Password: request.Password,
	}

	_, err := u.Save(c.DB)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"errors": map[string]interface{}{
				"email": err.Error(),
			},
		})
		return
	}

	ctx.JSON(http.StatusCreated, &u)
}

func (c AuthController) CheckAuth(ctx *gin.Context) {

}
