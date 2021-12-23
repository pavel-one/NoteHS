package routes

import (
	"app/controllers/AuthController"
)

func (r Route) Auth() {
	c := AuthController.New(r.DB)

	auth := r.Router.Group("auth")
	{
		auth.GET("check", c.CheckAuth)
		auth.POST("register", c.Register)
		auth.POST("auth", c.Auth)
	}
}
