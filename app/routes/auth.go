package routes

import (
	"app/controllers"
)

func (r Route) Auth() {
	userController := controllers.NewUser(r.DB)
	dialController := controllers.NewDialController(r.DB)

	user := r.Router.Group("user").Use(userController.AuthMiddleware)
	{
		user.GET("/", userController.User)
	}

	dial := r.Router.Group("dial").Use(dialController.AuthMiddleware)
	{
		dial.GET("/", dialController.GetAllDials)
		dial.PUT("/", dialController.CreateDial)

		dial.POST("/:id", dialController.EditDial)
	}
}
