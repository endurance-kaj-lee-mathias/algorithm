package metrics

import (
	"gitlab.com/kdg-ti/the-lab/teams-25-26/26-de-uitgeruste-it-ers/algorithm/internal/metrics/application"
	"gitlab.com/kdg-ti/the-lab/teams-25-26/26-de-uitgeruste-it-ers/algorithm/internal/metrics/transport"
)

func Wire() *transport.Handler {
	service := application.NewService()
	return transport.NewHandler(service)
}
