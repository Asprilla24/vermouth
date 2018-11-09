package api

import (
	"github.com/Asprilla24/vermouth/api/handlers"
	"github.com/Asprilla24/vermouth/database"
	"github.com/jinzhu/gorm"
)

type API struct {
	user *handlers.UserHandler
}

func New(db *gorm.DB) (*API, error) {
	userStore := database.NewUserStore(db)
	userHandler := handlers.NewUserHandler(userStore)

	api := &API{
		user: userHandler,
	}

	return api, nil
}
