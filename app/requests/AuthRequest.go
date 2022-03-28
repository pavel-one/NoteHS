package requests

import (
	"gopkg.in/guregu/null.v4"
)

type Auth struct {
	*BaseRequest
	Email    null.String `form:"email" json:"email" binding:"omitempty,email"`
	Password null.String `form:"password" json:"password" binding:"omitempty,min=6"`
	GoogleID null.String `form:"google_id" json:"google_id" binding:"omitempty"`
}

type Register struct {
	*BaseRequest
	Email    null.String `form:"email" json:"email" binding:"required,email"`
	Name     null.String `form:"name" json:"name" binding:"required"`
	Password null.String `form:"password" json:"password" binding:"required,min=6"`
}
