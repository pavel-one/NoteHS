package core

import (
	"app/base"
	"app/routes"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type App struct {
	DB     *base.DB
	Router base.Router
	Route  *routes.Route
}

func NewApp() *App {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	db := base.LoadDB()
	router := base.LoadRouter()
	route := routes.New(db, &router)

	return &App{
		DB:     db,
		Router: router,
		Route:  &route,
	}
}

func (a App) Run() {
	if os.Getenv("GIN_MODE") != "production" {
		err := a.Router.Run(":1200")
		if err != nil {
			log.Fatal("Ошибка запуска сервера " + err.Error())
		}
	} else {
		err := a.Router.Run("127.0.0.1:1200")
		if err != nil {
			log.Fatal("Ошибка запуска сервера " + err.Error())
		}
	}

}

func (a App) LoadRoutes(providers ...func()) {
	for _, provider := range providers {
		provider()
	}
}
