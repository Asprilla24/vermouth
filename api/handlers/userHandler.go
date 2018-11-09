package handlers

import "github.com/Asprilla24/vermouth/models"

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

func (handler *UserHandler) Router() {

}
