package controllers

import (
	"app/base"
	"app/helpers"
	"app/models"
	"app/requests"
	"app/resources"
	"github.com/gin-gonic/gin"
)

type DialController struct {
	*Controller
}

func NewDialController(db *base.DB) *DialController {
	controller := Controller{DB: db}

	return &DialController{&controller}
}

func (c DialController) GetAllDials(ctx *gin.Context) {
	token, _ := helpers.GetToken(ctx)
	user, _ := helpers.GetUserWithToken(token, c.DB)

	c.DB.Model(&user).Preload("Dials").First(&user)

	c.Success(resources.DialResources(user.Dials), ctx)
	return
}

func (c DialController) CreateDial(ctx *gin.Context) {
	var request requests.CreateDialRequest
	var dial models.Dial
	token, _ := helpers.GetToken(ctx)
	user, _ := helpers.GetUserWithToken(token, c.DB)

	if !requests.Validate(&request, ctx) {
		return
	}

	dial.FillWithRequest(c.DB, request)

	user.Dials = []models.Dial{dial}
	c.DB.Save(&user)
	dial = user.Dials[0] //With id

	dial.SetProcess(c.DB)
	go dial.CreateOrUpdateInfo(c.DB)

	c.Success(resources.DialResource(&dial), ctx)
	return
}
func (c DialController) EditDial(ctx *gin.Context) {
	var request requests.CreateDialRequest
	var dial models.Dial
	token, _ := helpers.GetToken(ctx)
	user, _ := helpers.GetUserWithToken(token, c.DB)

	c.DB.Where("id = ? and user_id = ?", ctx.Param("id"), user.ID).First(&dial)

	if dial.ID == 0 {
		c.Error(map[string]interface{}{
			"message": "Нет такого диала, или он вам не пренадлежит",
		}, ctx)
		return
	}

	if !requests.Validate(&request, ctx) {
		return
	}

	dial.FillWithRequest(c.DB, request)
	c.DB.Save(&dial)
	dial.SetProcess(c.DB)
	go dial.CreateOrUpdateInfo(c.DB)

	c.Success(resources.DialResource(&dial), ctx)
	return
}
