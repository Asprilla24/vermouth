package models

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

//UserModel for init gorm table and modeling
type UserModel struct {
	ID             int64     `gorm:"primary_key" json:"id"`
	Username       string    `gorm:"unique_index" json:"username"`
	Password       string    `gorm:"-" json:"password,omitemty"`
	HashedPassword string    `gorm:"size:100" json:"-"`
	CreatedAt      time.Time `gorm:"created_at" json:"created_at"`
}

//EncryptPassword set a new hashsed password to user
func (user *UserModel) EncryptPassword() {
	bcryptPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.HashedPassword = string(bcryptPassword[:])
}

func (user *UserModel) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.Must(uuid.NewV4()))
	scope.SetColumn("CreatedAt", time.Now())
	user.EncryptPassword()
	return nil
}

func (user *UserModel) TableName() string {
	return "users"
}
