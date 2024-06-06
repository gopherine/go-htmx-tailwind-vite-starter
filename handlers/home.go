package handlers

import (
	"net/http"

	"github.com/goemon/views/home"
	"github.com/labstack/echo/v4"
)

func Hello(c echo.Context) error {
	return Render(c, http.StatusOK, home.Index())
}
