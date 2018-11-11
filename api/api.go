package api

import (
	"strings"

	"github.com/Asprilla24/vermouth/api/handlers"
	"github.com/Asprilla24/vermouth/database"
	"github.com/Asprilla24/vermouth/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type API struct {
	echo        *echo.Echo
	userHandler *handlers.UserHandler
}

func New() *API {
	api := &API{
		echo: echo.New(),
	}

	api.InitializeHandler().InitializeRouter()

	return api
}

func (api *API) InitializeHandler() *API {
	db := database.InitializeDB(&models.UserModel{})

	userStore := database.NewUserStore(db)
	userHandler := handlers.NewUserHandler(userStore)

	api.userHandler = userHandler

	return api
}

func (api *API) InitializeRouter() *API {
	api.echo.Use(middleware.Logger())

	api.echo.POST("/login", api.userHandler.Login)
	api.echo.POST("/register", api.userHandler.CreateUser)

	userGroup := api.echo.Group("/user")
	api.userHandler.Router(userGroup)

	return api
}

func (api *API) Run(port string) {
	if !strings.Contains(port, ":") {
		port = ":" + port
	}

	api.echo.Logger.Fatal(api.echo.Start(port))
}
