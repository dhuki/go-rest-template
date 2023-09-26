package service

import (
	"context"
	"fmt"

	"github.com/dhuki/go-rest-template/internal/core/port"
	"github.com/dhuki/go-rest-template/pkg/logger"
)

type HealthService struct {
	repo port.HealthRepository
}

func NewHealthService(healthRepo port.HealthRepository) port.HealthService {
	return &HealthService{
		repo: healthRepo,
	}
}

func (h *HealthService) HealthCheck(ctx context.Context) (err error) {
	ctxName := fmt.Sprintf("%T.HealthCheck", h)

	if err = h.repo.Ping(ctx); err != nil {
		logger.Error(ctx, ctxName, "u.repository.Ping, got err: %v", err)
		return
	}

	return
}
