package requests

import (
	"errors"
	"github.com/gabriel-vasile/mimetype"
	"gopkg.in/guregu/null.v4"
	"mime/multipart"
	"strings"
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
	Image multipart.FileHeader `form:"image" json:"image" binding:"omitempty"`
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

func (r EditDialRequest) CheckUploadedFile() error {
	if r.Image.Size == 0 {
		return errors.New("not upload file")
	}

	image, _ := r.Image.Open()
	mime, _ := mimetype.DetectReader(image)
	allow := []string{"image/jpeg", "image/pjpeg", "image/png"}

	if !mimetype.EqualsAny(mime.String(), allow...) {
		return errors.New("Not allowed file, allow: " + strings.Join(allow, ", "))
	}

	return nil
}
