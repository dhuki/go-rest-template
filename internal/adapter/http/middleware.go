package http

import (
	"time"

	"github.com/dhuki/go-rest-template/pkg/logger"
	"github.com/labstack/echo/v4"
)

type requestLog struct {
	Timestamp     time.Time   `json:"timestamp"`
	CorrelationID interface{} `json:"correlationId"`
	Method        string      `json:"method"`
	URL           string      `json:"url"`
	Status        int         `json:"status"`
	ResponseTime  float64     `json:"responseTime"`
	ResponseSize  int64       `json:"responseSize"`
	ReqBody       interface{} `json:"requestBody"`
}

func LogMiddleware() echo.MiddlewareFunc {
	return echo.MiddlewareFunc(func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {
			ctx := c.Request().Context()
			start := time.Now()
			if err := next(c); err != nil {
				c.Error(err)
			}
			rl := requestLog{
				Timestamp:     start,
				CorrelationID: "",
				Method:        c.Request().Method,
				URL:           c.Request().URL.RequestURI(),
				Status:        c.Response().Status,
				ResponseTime:  time.Since(start).Seconds(),
				ResponseSize:  c.Response().Size,
			}
			logger.Info(ctx, "logger", "%+v", rl)

			return nil
		})
	})
}
