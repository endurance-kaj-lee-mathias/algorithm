package message

import (
	"gitlab.com/kdg-ti/the-lab/teams-25-26/26-de-uitgeruste-it-ers/algorithm/internal/application"
	"gitlab.com/kdg-ti/the-lab/teams-25-26/26-de-uitgeruste-it-ers/algorithm/internal/transport"
)

func Wire() *transport.Handler {
	service := application.NewService()
	return transport.NewHandler(service)
}
