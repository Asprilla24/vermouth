package database

import (
	"fmt"

	"github.com/Asprilla24/vermouth/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//ConnectDB to open connection based on configuration
func ConnectDB() (*gorm.DB, error) {
	config := config.GetConfig()

	dbURI := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Name,
		config.Database.Charset)

	db, err := gorm.Open(config.Database.Dialect, dbURI)
	if err != nil {
		return nil, err
	}

	return db, nil
}

//InitializeDB : for opening the Database and migrate an models
func InitializeDB(models ...interface{}) *gorm.DB {
	db, err := ConnectDB()
	if err != nil {
		panic(err.Error())
	}

	db.DB()
	for _, model := range models {
		db.AutoMigrate(model)
	}

	return db
}
