package config

//Config struct
type Config struct {
	DB *DBConfig
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

var dbConfig = DBConfig{
	Dialect:  "mysql",
	Username: "root",
	Password: "password",
	Host:     "127.0.0.1",
	Name:     "go_example",
	Charset:  DefaultCharset,
}

//GetConfig : get configuration
func GetConfig() *Config {
	return &Config{
		DB: &dbConfig,
	}
}
