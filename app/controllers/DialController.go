package controllers

import (
	"app/Services/ImageService"
	"app/base"
	"app/helpers"
	"app/models"
	"app/requests"
	"app/resources"
	"github.com/gin-gonic/gin"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
	"net/http"
	"time"
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

	c.DB.Model(&user).Preload("Dials", func(db *gorm.DB) *gorm.DB {
		db = db.Where("type = ? and deleted = ?", ctx.Query("type"), false).Order("created_at desc")
		return db
	}).First(&user)

	c.Success(resources.DialResources(user.Dials), ctx)
	return
}

func (c DialController) SyncPopularDials(ctx *gin.Context) {
	var request requests.SyncDialRequest
	var dials []models.Dial

	token, _ := helpers.GetToken(ctx)
	user, _ := helpers.GetUserWithToken(token, c.DB)

	if !requests.Validate(&request, ctx) {
		return
	}

	for _, v := range request.Dials {
		var dial models.Dial

		url := v.(map[string]interface{})["url"].(string)

		c.DB.Model(&dial).Where("user_id = ? and url = ? and type = ?", user.ID, url, 1).First(&dial)

		if dial.ID != 0 {
			dials = append(dials, dial)
			continue
		}

		dial.Url = url
		dial.UserID = user.ID
		dial.Type = 1

		c.DB.Save(&dial)

		dial.SetProcess(c.DB)
		go dial.CreateOrUpdateInfo(c.DB)
		dials = append(dials, dial)
	}

	c.Success(resources.DialResources(dials), ctx)
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
	var request requests.EditDialRequest
	var dial models.Dial
	token, _ := helpers.GetToken(ctx)
	user, _ := helpers.GetUserWithToken(token, c.DB)

	c.DB.Where("id = ? and user_id = ? and deleted = ?", ctx.Param("id"), user.ID, false).First(&dial)

	if dial.ID == 0 {
		c.Error(map[string]interface{}{
			"message": "Нет такого диала, или он вам не пренадлежит",
		}, ctx)
		return
	}

	updatePhoto := false

	if !requests.Validate(&request, ctx) {
		return
	}

	err := request.CheckUploadedFile() //Validate mime file

	if request.Image.Size > 0 {
		if err != nil {
			c.Error(map[string]interface{}{
				"image": err.Error(),
			}, ctx)

			return
		}
	}

	if request.Url != dial.Url {
		updatePhoto = true
	}

	dial.FillWithRequest(c.DB, request) //Fill simple field

	//If image exists
	if request.Image.Size > 0 {
		imagePath, err := ImageService.SaveImageWithForm(request.Image, dial.Url, dial.UserID)

		if err != nil {
			c.Error(map[string]interface{}{
				"image": "Error processing image: " + err.Error(),
			}, ctx)

			return
		}

		dial.Screen = null.StringFrom(imagePath)
	}

	c.DB.Save(&dial)

	//Try create site screenshot after
	defer func() {
		if updatePhoto && (request.Image.Size == 0) {
			dial.SetProcess(c.DB)
			go dial.UpdatePhoto(c.DB)
		}
	}()

	c.Success(resources.DialResource(&dial), ctx)
	return
}

func (c DialController) GetDialInfo(ctx *gin.Context) {
	var dial models.Dial
	token, _ := helpers.GetToken(ctx)
	user, _ := helpers.GetUserWithToken(token, c.DB)

	c.DB.Where("id = ? and user_id = ? and deleted = ?", ctx.Param("id"), user.ID, false).First(&dial)

	if dial.ID == 0 {
		c.Error(map[string]interface{}{
			"message": "Нет такого диала, или он вам не пренадлежит",
		}, ctx)
		return
	}

	if !dial.Final {
		diff := time.Now().Sub(dial.UpdatedAt).Minutes()
		if diff > 1 {
			dial.SetProcessEnd(c.DB)
		}
	}

	c.Success(resources.DialResource(&dial), ctx)
	return
}

func (c DialController) DropDial(ctx *gin.Context) {
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

	fake := false
	if dial.Type == 1 {
		fake = true
	}

	dial.DropDialWithFiles(c.DB, fake)

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
	})
	return
}
