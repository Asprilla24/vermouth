package handlers

import (
	"time"

	"github.com/Asprilla24/vermouth/config"
	"github.com/Asprilla24/vermouth/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang.org/x/crypto/bcrypt"
)

//UserStore defines database operation for user
type UserStore interface {
	CreateUser(user *models.UserModel) error
	GetUser(username string) (*models.UserModel, error)
	GetUsers() (*[]models.UserModel, error)
}

//UserHandler implement user management handler
type UserHandler struct {
	store     UserStore
	tokenCode string
}

//NewUserHandler create and return user handler
func NewUserHandler(store UserStore) *UserHandler {
	config := config.GetConfig()

	return &UserHandler{
		store:     store,
		tokenCode: config.TokenCode,
	}
}

//Router create and return router for UserHandler
func (handler *UserHandler) Router(g *echo.Group) {
	g.Use(middleware.JWT([]byte(handler.tokenCode)))
	g.GET("/GetAll", handler.GetUsers)
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

func (handler *UserHandler) GetUsers(e echo.Context) error {
	result, err := handler.store.GetUsers()
	if err != nil {
		return ShowErrorResponse(e, err)
	}

	return ShowSuccessResponse(e, result)
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

	println(username + ", " + handler.tokenCode)

	// Create jwt token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = username
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response
	t, err := token.SignedString([]byte(handler.tokenCode))
	if err != nil {
		ShowErrorResponse(e, err)
	}

	return ShowSuccessResponse(e, t)
}
