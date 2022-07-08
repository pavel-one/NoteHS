package controllers

import (
	"app/base"
	"app/helpers"
	"app/models"
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
		var post models.Post

		c.DB.Where("user_id = ? and id = ?", user.ID, request.PostId).Model(&post)

		//TODO: Проверять существование такого поста
		user.Settings.PostId = request.PostId
	}

	user.Settings.Component = request.Component

	c.DB.Save(&user)
	c.DB.Save(user.Settings)

	resource := resources.GetSettingResource(user.Settings)

	c.Success(resource, ctx)
}

func NewUser(db *base.DB) *UserController {
	controller := Controller{DB: db}

	return &UserController{&controller}
}
