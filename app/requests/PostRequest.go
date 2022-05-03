package requests

import (
	"gopkg.in/guregu/null.v4"
)

type PostRequest struct {
	*BaseRequest
	Id          null.String `json:"id" binding:"omitempty,uuid"`
	Name        string      `json:"name" binding:"required,max=255"`
	Description null.String `json:"description" binding:"omitempty,max=255"`
	Data        struct {
		Time    int64  `json:"time"  binding:"-"`
		Version string `json:"version"  binding:"-"`
		Blocks  []map[string]interface{}
	} `json:"data" binding:"-"`
}
