package resources

import (
	"app/models"
	"gopkg.in/guregu/null.v4"
)

type UserResource struct {
	Id    uint        `json:"id"`
	Name  null.String `json:"name"`
	Email null.String `json:"email"`
	Token null.String `json:"token"`
}

func GetUserResource(user *models.User) UserResource {
	resource := UserResource{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	if user.Token.Token.Valid {
		resource.Token = user.Token.Token
	}

	return resource
}
