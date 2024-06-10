package handlers

import (
	"net/http"

	"github.com/golang_starter/views/pages/home"
	"github.com/labstack/echo/v4"
)

func Hello(c echo.Context) error {
	return Render(c, http.StatusOK, home.Index())
}
