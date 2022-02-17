package requests

type CreateDialRequest struct {
	Url         string `form:"url" json:"url" binding:"exists,url"`
	Name        string `form:"name" json:"name"`
	Description string `form:"description" json:"description"`
}
