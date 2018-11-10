package handlers

import (
	"github.com/Asprilla24/vermouth/models"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

//UserStore defines database operation for user
type UserStore interface {
	CreateUser(user *models.UserModel) error
	GetUser(username string) (*models.UserModel, error)
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
	g.POST("/login", handler.Login)
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

func (handler *UserHandler) Login(e echo.Context) error {
	var (
		username, password string
	)

	username = e.FormValue("username")
	password = e.FormValue("password")

	user := &models.UserModel{}
	user, err := handler.store.GetUser(username)
	if err != nil {
		return ShowErrorResponse(e, err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	if err != nil {
		return ShowErrorResponse(e, err)
	}

	return ShowSuccessResponse(e, "")
}
