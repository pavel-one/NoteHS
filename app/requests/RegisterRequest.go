package requests

import "gopkg.in/guregu/null.v4"

type Register struct {
	Email    null.String `form:"email" json:"email" binding:"required,email"`
	Name     null.String `form:"name" json:"name" binding:"required"`
	Password null.String `form:"password" json:"password" binding:"required,min=6"`
}
