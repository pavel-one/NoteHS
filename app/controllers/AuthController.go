package controllers

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
	*Controller
}

func NewAuth(db *base.DB) *AuthController {
	controller := Controller{DB: db}

	return &AuthController{&controller}
}

func (c AuthController) Auth(ctx *gin.Context) {
	var request requests.Auth
	var user models.User

	if err := ctx.ShouldBindJSON(&request); err != nil {
		e := ValidationExeption.New(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": e.FormatToFront()})
		return
	}

	c.DB.Model(&user).Where("email = ?", request.Email).First(&user)

	if user.ID == 0 {
		c.Error(map[string]interface{}{
			"email": "Не правильное имя пользователя или пароль",
		}, ctx)
		return
	}

	if !user.CheckPasswordHash(request.Password) {
		c.Error(map[string]interface{}{
			"email": "Не правильное имя пользователя или пароль",
		}, ctx)
		return
	}

	resource := resources.GetUserResource(&user)

	c.Success(resource, ctx)
	return
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
		c.Error(map[string]interface{}{
			"email": err.Error(),
		}, ctx)
		return
	}

	resource := resources.GetUserResource(&u)

	c.Success(resource, ctx)
}

func (c AuthController) CheckAuth(ctx *gin.Context) {

}
