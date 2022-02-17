package requests

type Register struct {
	Email    string `form:"email" json:"email" binding:"exists,email"`
	Name     string `form:"name" json:"name" binding:"exists"`
	Password string `form:"password" json:"password" binding:"exists,min=6"`
}
