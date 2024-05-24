package api

import (
	"github.com/labstack/echo/v4"
)

func (a *App) Greetings(ctx echo.Context) error {
	return ctx.HTML(200, greetingsHtmx())
}

// func (a *App) Test(c *fiber.Ctx) error {
// 	c.Type()
// }
