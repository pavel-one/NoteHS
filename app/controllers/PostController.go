package controllers

import (
	"app/base"
	"app/helpers"
	"app/models"
	"app/requests"
	"app/resources"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
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
	token, _ := helpers.GetToken(ctx)
	user, _ := helpers.GetUserWithToken(token, c.DB)

	var post models.Post
	var request requests.PostRequest

	if !requests.Validate(&request, ctx) {
		return
	}

	id, err := uuid.NewV4()
	if err != nil {
		return
	}

	post.Uuid = id.String()
	post.Name = request.Name
	post.Description = request.Description
	post.Public = false
	post.UserId = user.ID
	post.PostData = request.Data.ToString()

	c.DB.Create(&post)

	c.Success(resources.PostResource(&post), ctx)

	return
}

func (c PostController) Remove(ctx *gin.Context) {
	token, _ := helpers.GetToken(ctx)
	user, _ := helpers.GetUserWithToken(token, c.DB)

	var post models.Post
	c.DB.Where("user_id = ? and uuid = ?", user.ID, ctx.Param("id")).First(&post)

	if post.Uuid == "" {
		c.Error(map[string]interface{}{
			"id": "Не найдена такая заметка",
		}, ctx)

		return
	}

	c.DB.Where("user_id = ? and uuid = ?", user.ID, ctx.Param("id")).Delete(&post)
	c.Success(resources.PostResource(&post), ctx)

	return
}
