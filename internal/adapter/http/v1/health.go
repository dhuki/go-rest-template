package v1

import (
	"net/http"

	httpHelper "github.com/dhuki/go-rest-template/internal/adapter/http"
	"github.com/labstack/echo/v4"
)

func (s *svc) RegistHealthRoute(app *echo.Group) {
	v1GroupHealth := app.Group("/health")

	v1GroupHealth.GET("", s.healthCheck())
}

func (s *svc) healthCheck() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		if err := s.handler.HealthService.HealthCheck(ctx); err != nil {
			return httpHelper.ResponseError(c, http.StatusInternalServerError, err)
		}

		return httpHelper.ResponseSuccess(c, "success", nil)
	}
}
