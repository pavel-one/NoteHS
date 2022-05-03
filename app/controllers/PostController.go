package controllers

import (
	"app/base"
	"app/helpers"
	"app/models"
	"app/requests"
	"app/resources"
	"fmt"
	"github.com/gin-gonic/gin"
)

type PostController struct {
	*Controller
}

func NewPostController(db *base.DB) *PostController {
	controller := Controller{DB: db}

	return &PostController{&controller}
}

func (c PostController) All(ctx *gin.Context) {
	token, _ := helpers.GetToken(ctx)
	user, _ := helpers.GetUserWithToken(token, c.DB)

	var posts []models.Post

	c.DB.Where("user_id = ?", user.ID).Order("updated_at desc").Find(&posts)

	c.Success(resources.PostResources(posts), ctx)
	return
}

func (c PostController) UpdateOrCreate(ctx *gin.Context) {
	//token, _ := helpers.GetToken(ctx)
	//user, _ := helpers.GetUserWithToken(token, c.DB)

	//var post models.Post
	var request requests.PostRequest

	if !requests.Validate(&request, ctx) {
		return
	}

	fmt.Println(request)
}
