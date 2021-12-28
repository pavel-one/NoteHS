package resources

import (
	"app/models"
	"os"
	"strconv"
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

	ttl, _ := strconv.Atoi(os.Getenv("TOKEN_TTL_DAY"))

	if user.Token.Token != "" {
		resource.Token = map[string]interface{}{
			"token":  user.Token.Token,
			"ttlDay": user.Token.CreatedAt.AddDate(0, 0, ttl).Sub(user.Token.CreatedAt).Seconds() / 86400,
		}
	}

	return resource
}
