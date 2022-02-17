package requests

import (
	"gopkg.in/guregu/null.v4"
)

type Auth struct {
	Email    null.String `form:"email" json:"email" binding:"email"`
	Password null.String `form:"password" json:"password" binding:"min=6"`
	GoogleID null.String `form:"google_id" json:"google_id" binding:"required,uuid"`
}
