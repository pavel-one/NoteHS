package controllers

import (
	"app/Services/Scrapper"
	"app/base"
	"app/helpers"
	"app/models"
	"app/requests"
	"app/resources"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
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
	token, _ := helpers.GetToken(ctx)
	user, _ := helpers.GetUserWithToken(token, c.DB)

	if !requests.Validate(&request, ctx) {
		return
	}

	dial := models.Dial{
		Url: request.Url,
	}

	user.Dials = []models.Dial{dial}
	c.DB.Save(&user)
	dial = user.Dials[0]

	//TODO: Вынести нахер отсюда
	go func() {
		defer func() {
			dial.Final = true
			c.DB.Save(&dial)
		}()

		url, err := Scrapper.GetUrlInfo(dial.Url, strconv.Itoa(int(dial.ID)))
		if err != nil {
			log.Println(err)
			return
		}
		dial.Name = url.Title
		dial.Screen = url.Screen

	}()

	c.Success(resources.DialResource(&dial), ctx)
	return
}
