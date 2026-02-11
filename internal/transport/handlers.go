package transport

import (
	"encoding/json"
	"net/http"

	"gitlab.com/kdg-ti/the-lab/teams-25-26/26-de-uitgeruste-it-ers/algorithm/internal/domain"
	"gitlab.com/kdg-ti/the-lab/teams-25-26/26-de-uitgeruste-it-ers/algorithm/internal/response"
	"gitlab.com/kdg-ti/the-lab/teams-25-26/26-de-uitgeruste-it-ers/algorithm/internal/transport/models"
)

func (h *Handler) CreateMessage(w http.ResponseWriter, r *http.Request) {
	var m models.MessageModel
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		response.WriteError(w, http.StatusBadRequest, InvalidQuery)
		return
	}

	if m.HeartRate <= 0 && m.SleepTimeMinutes <= 0 && m.BloodPressure == "" {
		response.WriteError(w, http.StatusBadRequest, InvalidQuery)
		return
	}

	metric := domain.HealthMetric{
		HeartRate:        m.HeartRate,
		SleepTimeMinutes: m.SleepTimeMinutes,
		BloodPressure:    m.BloodPressure,
	}

	stress, err := h.service.ComputeStress(r.Context(), metric)
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	response.Write(w, http.StatusOK, models.StressModel{Stress: stress})
}
