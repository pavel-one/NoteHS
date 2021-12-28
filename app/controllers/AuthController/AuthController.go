package AuthController

import (
	"app/base"
	"app/exceptions/ValidationExeption"
	"app/models"
	"app/requests"
	"app/resources"
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
		Name:     request.Name,
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

	resource := resources.GetUserResource(&u)

	ctx.JSON(http.StatusCreated, &resource)
}

func (c AuthController) CheckAuth(ctx *gin.Context) {

}
