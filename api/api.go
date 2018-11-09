package api

import (
	"github.com/Asprilla24/vermouth/api/handlers"
	"github.com/Asprilla24/vermouth/database"
	"github.com/Asprilla24/vermouth/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type API struct {
	user *handlers.UserHandler
}

func New() *API {
	db := database.InitializeDB(&models.UserModel{})

	userStore := database.NewUserStore(db)
	userHandler := handlers.NewUserHandler(userStore)

	api := &API{
		user: userHandler,
	}

	return api
}

func (api *API) Run(port string) {
	e := echo.New()
	e.Use(middleware.Logger())

	userHandler := e.Group("/user")

	api.user.Router(userHandler)

	e.Logger.Fatal(e.Start(port))
}
