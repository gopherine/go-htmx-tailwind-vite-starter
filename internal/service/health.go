package service

import (
	"context"

	"github.com/golang_starter/internal/logger"
)

type HealthService struct{}

func NewHealthService() *HealthService {
	return &HealthService{}
}

func (s *HealthService) CheckStatus(ctx context.Context) map[string]string {
	logger.Debug(ctx, "Checking service status")
	return map[string]string{"status": "ok"}
}
