package requests

type CreateDialRequest struct {
	Url string `form:"url" json:"url" binding:"required,url"`
}
