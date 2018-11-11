# Go Restful API Boilerplate

Easily extendible RESTful API boilerplate aiming to follow idiomatic go and best practice.

The goal of this boiler is to have a solid and structured foundation to build upon on.

### Features
The following feature set is a minimal selection of typical Web API requirements:

- CLI features using [cobra](https://github.com/spf13/cobra)
- ORM support including migrations using [gorm](https://github.com/jinzhu/gorm)
- Base APP with [echo](https://github.com/labstack/echo) and middleware
- JWT Authentication using [jwt-go](https://github.com/dgrijalva/jwt-go)

### Start Application
- Clone this repository
- Setting your Database and Token secret key in config
- Build the application: ```go build``` to create ```vermouth``` binary or use ```go run main.go``` instead in the following commands
- Run ```vermouth``` for cobra generated help message.
- Run the application: ```vermouth serve``` the default port is 8080
- Or ```vermouth serve -p [port]```
- Or ```vermouth serve --port [port]```

### Contributing

Any feedback are welcome and highly appreciated.
