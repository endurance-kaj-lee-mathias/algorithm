package application

import (
	"context"

	"gitlab.com/kdg-ti/the-lab/teams-25-26/26-de-uitgeruste-it-ers/algorithm/internal/metrics/domain"
)

type Service interface {
	ComputeStress(ctx context.Context, metric domain.HealthMetric) (int, error)
}
