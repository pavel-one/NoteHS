package routes

import (
	"app/controllers"
)

func (r *Route) Api() {
	var userController = controllers.NewUser(r.DB)

	var api = r.Router.Group("api")
	{
		api.GET("/", userController.Index)
		api.GET("/create", userController.Create)
	}
}
