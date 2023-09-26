package port

import "context"

// port input
type HealthService interface {
	HealthCheck(ctx context.Context) (err error)
}

// port output
type HealthRepository interface {
	Ping(ctx context.Context) (err error)
}
