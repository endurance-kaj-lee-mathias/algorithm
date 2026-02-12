package health

import (
	"net/http"

	"gitlab.com/kdg-ti/the-lab/teams-25-26/26-de-uitgeruste-it-ers/algorithm/internal/metrics/response"
)

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	status := Status{
		Backend: "UP",
	}

	response.Write(w, http.StatusOK, status)
}
