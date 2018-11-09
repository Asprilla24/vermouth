package models

import (
	"time"
)

//UserModel for init gorm table and modeling
type UserModel struct {
	ID             string    `gorm:"size:100" json:"id"`
	Username       string    `gorm:"username" json:"username"`
	Password       string    `gorm:"-" json:"password"`
	PasswordHashed string    `gorm:"size:100" json:"-"`
	CreatedAt      time.Time `gorm:"created_at" json:"created_at"`
}
