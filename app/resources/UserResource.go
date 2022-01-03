package resources

import (
	"app/models"
)

type UserResource struct {
	Id    uint                   `json:"id"`
	Name  string                 `json:"name"`
	Email string                 `json:"email"`
	Token map[string]interface{} `json:"token"`
}

func GetUserResource(user *models.User) UserResource {
	resource := UserResource{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	if user.Token.Token != "" {
		resource.Token = map[string]interface{}{
			"token": user.Token.Token,
		}
	}

	return resource
}
