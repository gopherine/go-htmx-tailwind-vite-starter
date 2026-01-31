package routes

import "github.com/labstack/echo/v4"

// RegisterAll calls generated and custom routes.
func RegisterAll(e *echo.Echo) {
	Register(e)
	RegisterCustom(e)
}

// RegisterCustom is the place to wire manual routes that include custom logic.
func RegisterCustom(e *echo.Echo) {
}
