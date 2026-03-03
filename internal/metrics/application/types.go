package application

import (
	"context"

	"gitlab.com/kdg-ti/the-lab/teams-25-26/26-de-uitgeruste-it-ers/algorithm/internal/metrics/domain"
)

type Service interface {
	ComputeStress(ctx context.Context, userID string, samples []domain.Sample) (domain.StressResult, error)
}

type service struct{}

func NewService() Service {
	return &service{}
}
