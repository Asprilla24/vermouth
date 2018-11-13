package config

import (
	"fmt"
	"reflect"

	"github.com/spf13/viper"
)

//Config struct
type Config struct {
	Database  *DBConfig
	TokenCode string
}

//DBConfig for database configuration
type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Host     string
	Name     string
	Charset  string
}

//DefaultCharset : utf8
const DefaultCharset = "utf8"
const DefaultDialect = "mysql"
const DefaultTokenCode = "mySecretCode"

var configuration = &Config{}

func InitializeConfig() {
	viper := viper.New()
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	config := &Config{}

	if err := viper.ReadInConfig(); err != nil {
		panic(err.Error())
	}

	if err := viper.Unmarshal(config); err != nil {
		panic(err.Error())
	}

	fmt.Println(config)

	configuration = config.SetDefaultValueIfEmpty()
}

//GetConfig : get configuration
func GetConfig() *Config {
	return configuration
}

func (config *Config) SetDefaultValueIfEmpty() *Config {
	if IsZeroOfUnderlyingType(config.TokenCode) {
		config.TokenCode = DefaultTokenCode
	}

	if IsZeroOfUnderlyingType(config.Database.Charset) {
		config.Database.Charset = DefaultCharset
	}

	if IsZeroOfUnderlyingType(config.Database.Dialect) {
		config.Database.Dialect = DefaultDialect
	}

	return config
}

func IsZeroOfUnderlyingType(x interface{}) bool {
	return reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}
