package handlers

import (
	"github.com/Asprilla24/vermouth/auth"
	"github.com/Asprilla24/vermouth/models"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

type MainHandler struct {
	userStore UserStore
}

func NewMainHandler(store UserStore) *MainHandler {
	return &MainHandler{
		userStore: store,
	}
}

func (handler *MainHandler) Router(g *echo.Group) {
	g.POST("/login", handler.Login)
	g.POST("/register", handler.Register)
}

func (handler *MainHandler) Login(e echo.Context) error {
	req := &models.UserModel{}
	if err := e.Bind(req); err != nil {
		return ShowErrorResponse(e, err)
	}

	user, err := handler.userStore.GetUser(req.Username)
	if err != nil {
		return ShowErrorResponse(e, err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(req.Password))
	if err != nil {
		return ShowErrorResponse(e, err)
	}

	token, err := auth.CreateJWTToken(user.Username, false)
	if err != nil {
		return ShowErrorResponse(e, err)
	}

	return ShowSuccessResponse(e, token)
}

func (handler *MainHandler) Register(e echo.Context) error {
	u := models.UserModel{}
	if err := e.Bind(&u); err != nil {
		return ShowErrorResponse(e, err)
	}

	if err := handler.userStore.CreateUser(&u); err != nil {
		return ShowErrorResponse(e, err)
	}

	return ShowSuccessResponse(e, "")
}
