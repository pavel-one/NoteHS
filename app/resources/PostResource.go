package resources

import (
	"app/Services/Human"
	"app/models"
	"gopkg.in/guregu/null.v4"
)

type postResource struct {
	Id          string      `json:"id"`
	Data        string      `json:"data"`
	Name        string      `json:"name"`
	Description null.String `json:"description"`
	Public      bool        `json:"public"`
	UpdatedAt   string      `json:"date"`
}

func PostResource(post *models.Post) *postResource {
	resource := postResource{
		Id:          post.Uuid,
		Data:        post.PostData,
		Name:        post.Name,
		Description: post.Description,
		Public:      post.Public,
		UpdatedAt:   Human.Time(post.UpdatedAt),
	}

	return &resource
}

func PostResources(posts []models.Post) []*postResource {
	var out []*postResource
	for _, post := range posts {
		out = append(out, PostResource(&post))
	}

	return out
}
