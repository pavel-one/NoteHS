package controllers

import (
	"app/base"
	"app/helpers"
	"app/requests"
	"app/resources"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	*Controller
}

func (c UserController) User(ctx *gin.Context) {
	token, _ := helpers.GetToken(ctx)
	user, _ := helpers.GetUserWithToken(token, c.DB)
	resource := resources.GetUserResource(&user)

	c.Success(resource, ctx)
}

func (c UserController) SetSetting(ctx *gin.Context) {
	token, _ := helpers.GetToken(ctx)
	user, _ := helpers.GetUserWithToken(token, c.DB)
	var request requests.SetSettingRequest

	if !requests.Validate(&request, ctx) {
		return
	}

	if user.Settings == nil {
		user.SetDefaultSettings()
	}

	if request.PostId != "" {
		user.Settings.PostId = request.PostId
	}

	user.Settings.Component = request.Component

	c.DB.Save(&user)

	resource := resources.GetSettingResource(user.Settings)

	c.Success(resource, ctx)
}

func NewUser(db *base.DB) *UserController {
	controller := Controller{DB: db}

	return &UserController{&controller}
}
