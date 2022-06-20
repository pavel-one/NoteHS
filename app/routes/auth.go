package routes

import (
	"app/controllers"
)

func (r Route) Auth() {
	userController := controllers.NewUser(r.DB)
	dialController := controllers.NewDialController(r.DB)
	postController := controllers.NewPostController(r.DB)

	user := r.Router.Group("user").Use(userController.AuthMiddleware)
	{
		user.GET("/", userController.User)
		user.POST("/sync/popular", dialController.SyncPopularDials)
		user.POST("/settings", userController.SetSetting)
	}

	dial := r.Router.Group("dial").Use(dialController.AuthMiddleware)
	{
		dial.GET("/", dialController.GetAllDials)
		dial.PUT("/", dialController.CreateDial)

		dial.GET("/:id", dialController.GetDialInfo)
		dial.POST("/:id", dialController.EditDial)
		dial.DELETE("/:id", dialController.DropDial)
	}

	post := r.Router.Group("posts").Use(postController.AuthMiddleware)
	{
		post.GET("/", postController.All)
		post.POST("/", postController.UpdateOrCreate)
		post.DELETE("/:id", postController.Remove)
	}
}
