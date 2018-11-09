package handlers

import (
	"github.com/Asprilla24/vermouth/models"
	"github.com/labstack/echo"
)

//UserStore defines database operation for user
type UserStore interface {
	CreateUser(user *models.UserModel) error
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
	g.POST("", handler.CreateUser)
}

func (handler *UserHandler) CreateUser(e echo.Context) error {
	u := models.UserModel{}
	if err := e.Bind(&u); err != nil {
		return ShowErrorResponse(e, err)
	}

	if err := handler.store.CreateUser(&u); err != nil {
		return ShowErrorResponse(e, err)
	}

	return ShowSuccessResponse(e, "")
}
