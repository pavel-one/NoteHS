package requests

import "gopkg.in/guregu/null.v4"

type CreateDialRequest struct {
	Url         string      `form:"url" json:"url" binding:"required,url"`
	Name        null.String `form:"name" json:"name" binding:"omitempty,max=255"`
	Description null.String `form:"description" json:"description" binding:"omitempty,max=255"`
}
