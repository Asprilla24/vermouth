package handlers

import (
	"github.com/Asprilla24/vermouth/config"

	"github.com/Asprilla24/vermouth/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//UserStore defines database operation for user
type UserStore interface {
	CreateUser(user *models.UserModel) error
	GetUser(username string) (*models.UserModel, error)
	GetUsers() (*[]models.UserModel, error)
}

//UserHandler implement user management handler
type UserHandler struct {
	store UserStore
}

//NewUserHandler create and return user handler
func NewUserHandler(store UserStore) *UserHandler {
	return &UserHandler{
		store: store,
	}
}

//Router create and return router for UserHandler
func (handler *UserHandler) Router(g *echo.Group) {
	g.Use(middleware.JWT([]byte(config.GetConfig().TokenCode)))
	g.GET("/GetAll", handler.GetUsers)
}

func (handler *UserHandler) GetUsers(e echo.Context) error {
	result, err := handler.store.GetUsers()
	if err != nil {
		return ShowErrorResponse(e, err)
	}

	return ShowSuccessResponse(e, result)
}
