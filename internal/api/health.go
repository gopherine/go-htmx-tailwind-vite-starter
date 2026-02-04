package api

import (
	"github.com/golang_starter/internal/logger"
	"github.com/golang_starter/internal/service"
	"github.com/labstack/echo/v4"
)

type HealthHandler struct {
	service *service.HealthService
}

func NewHealthHandler(service *service.HealthService) *HealthHandler {
	return &HealthHandler{service: service}
}

func (h *HealthHandler) Check(c echo.Context) error {
	ctx := logger.WithSpan(c.Request().Context(), "API")
	logger.Debug(ctx, "Health check requested")

	status := h.service.CheckStatus(ctx)
	return c.JSON(200, status)
}
