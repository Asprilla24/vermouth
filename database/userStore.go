package database

import (
	"github.com/Asprilla24/vermouth/models"
	"github.com/jinzhu/gorm"
)

//UserStore implement database operation for user management
type UserStore struct {
	db *gorm.DB
}

//NewUserStore : Return an UserStore
func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (s *UserStore) CreateUser(user *models.UserModel) error {
	err := s.db.Create(&user).Error
	return err
}
