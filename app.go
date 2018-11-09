package vermouth

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"github.com/labstack/echo"
)

//App struct for vermouth
type App struct {
	echo *echo.Echo
}

//New : for create a new Vermouth App with database configuration
func New() *App {
	app := new(App)
	app.echo = echo.New()
	return app
}

//InitializeDB : for Initialize DB with config and models to migrate
func (app *App) InitializeDB(config Config, models Models) *gorm.DB {
	return InitializeDB(config, models)
}

//Run : for running vermouth on what port
func (app *App) Run(port string) {
	app.echo.Logger.Fatal(app.echo.Start(port))
}

//SetRouters : for set routers
func (app *App) SetRouters(routers []Router) {
	app.Iris.Use(ShowRequestedAPI)
	for _, router := range routers {
		app.Iris.Handle(router.Method, router.Path, router.Handler)
	}
}

//ShowRequestedAPI : for logging requested API from command promt
func ShowRequestedAPI(ctx iris.Context) {
	ctx.Application().Logger().Printf("%s:	%s\n", ctx.Method(), ctx.Path())
	ctx.Next()
}
