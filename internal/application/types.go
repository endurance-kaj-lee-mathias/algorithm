package application

import (
	"context"

	"gitlab.com/kdg-ti/the-lab/teams-25-26/26-de-uitgeruste-it-ers/algorithm/internal/domain"
)

type Service interface {
	Save(ctx context.Context, msg domain.HealthMetric) error
	GetAll(ctx context.Context) ([]domain.HealthMetric, error)
}
