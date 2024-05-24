package api

import "github.com/labstack/echo/v4"

func (a *App) Routers() {
	a.e.GET("/", a.Greetings)
	// api := a.e.Group("/api")

}

type App struct {
	e    *echo.Echo
	port string
}

func NewApp(port string) *App {
	return &App{
		e:    echo.New(),
		port: port,
	}
}
func (a *App) Run() {
	a.Routers()
	a.e.Logger.Fatal(a.e.Start(":" + a.port))
}
