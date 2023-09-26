package http

import (
	"github.com/dhuki/go-rest-template/internal/adapter/repository"
	"github.com/dhuki/go-rest-template/internal/core/domain/service"
	"github.com/dhuki/go-rest-template/internal/core/port"
)

type Handler struct {
	HealthService port.HealthService
}

func NewHandler(repo repository.IRepository) Handler {
	return Handler{
		HealthService: service.NewHealthService(repo),
	}
}
