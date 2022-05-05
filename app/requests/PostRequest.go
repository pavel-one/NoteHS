package requests

import (
	"app/types"
	"gopkg.in/guregu/null.v4"
)

type PostRequest struct {
	*BaseRequest
	Id          null.String    `json:"id" binding:"omitempty,uuid"`
	Name        string         `json:"name" binding:"required,max=255"`
	Description null.String    `json:"description" binding:"omitempty,max=255"`
	Data        types.PostData `json:"data" binding:"-"`
}
