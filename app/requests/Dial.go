package requests

import (
	"gopkg.in/guregu/null.v4"
	"mime/multipart"
)

type DialRequest struct {
	*BaseRequest
	Url         string      `form:"url" json:"url" binding:"required,url"`
	Name        null.String `form:"name" json:"name" binding:"omitempty,max=255"`
	Description null.String `form:"description" json:"description" binding:"omitempty,max=255"`
}

type CreateDialRequest struct {
	*DialRequest
}

type EditDialRequest struct {
	*DialRequest
	Image multipart.File `form:"image" json:"image" binding:"omitempty"`
}

func (r DialRequest) GetName() null.String {
	return r.Name
}

func (r DialRequest) GetDescription() null.String {
	return r.Description
}

func (r DialRequest) GetUrl() string {
	return r.Url
}
