//go:generate go run ../../tools/generate_routes
package routes

import (
	"github.com/golang_starter/internal/api"
	"github.com/golang_starter/internal/service"
	"github.com/labstack/echo/v4"
)

// RegisterAll calls generated and custom routes.
func RegisterAll(e *echo.Echo) {
	RegisterUIRoutes(e)
	RegisterAPIRoutes(e)
}

// RegisterUIRoutes registers auto-generated page routes from templ files.
func RegisterUIRoutes(e *echo.Echo) {
	Register(e)
}

// RegisterAPIRoutes registers API routes.
func RegisterAPIRoutes(e *echo.Echo) {
	healthService := service.NewHealthService()
	healthHandler := api.NewHealthHandler(healthService)

	e.GET("/health", healthHandler.Check)
}
