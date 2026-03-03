package transport

import (
	"errors"
	"net/http"

	"gitlab.com/kdg-ti/the-lab/teams-25-26/26-de-uitgeruste-it-ers/algorithm/internal/metrics/application"
	"gitlab.com/kdg-ti/the-lab/teams-25-26/26-de-uitgeruste-it-ers/algorithm/internal/metrics/domain"
	"gitlab.com/kdg-ti/the-lab/teams-25-26/26-de-uitgeruste-it-ers/algorithm/internal/metrics/request"
	"gitlab.com/kdg-ti/the-lab/teams-25-26/26-de-uitgeruste-it-ers/algorithm/internal/metrics/response"
	"gitlab.com/kdg-ti/the-lab/teams-25-26/26-de-uitgeruste-it-ers/algorithm/internal/metrics/transport/models"
)

func (h *Handler) ComputeStress(w http.ResponseWriter, r *http.Request) {
	var body models.ComputeRequest
	if err := request.Decode(r, &body); err != nil {
		response.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if body.UserID == "" {
		response.WriteError(w, http.StatusBadRequest, errors.New("user_id is required"))
		return
	}

	if len(body.Samples) == 0 {
		response.WriteError(w, http.StatusBadRequest, errors.New("samples must not be empty"))
		return
	}

	samples := make([]domain.Sample, len(body.Samples))
	for i, s := range body.Samples {
		if s.HeartRate <= 0 {
			response.WriteError(w, http.StatusBadRequest, errors.New("each sample must have a valid heart_rate"))
			return
		}
		if s.RMSSD <= 0 {
			response.WriteError(w, http.StatusBadRequest, errors.New("each sample must have a valid rmssd"))
			return
		}
		if s.RestingHeartRate <= 0 {
			response.WriteError(w, http.StatusBadRequest, errors.New("each sample must have a valid resting_heart_rate"))
			return
		}
		samples[i] = domain.Sample{
			HeartRate:        s.HeartRate,
			RMSSD:            s.RMSSD,
			RestingHeartRate: s.RestingHeartRate,
			Steps:            s.Steps,
			SleepDebtHours:   s.SleepDebtHours,
			RecordedAt:       s.RecordedAt,
		}
	}

	result, err := h.service.ComputeStress(r.Context(), body.UserID, samples)
	if err != nil {
		if errors.Is(err, application.ErrNotEnoughSamples) {
			response.WriteError(w, http.StatusUnprocessableEntity, err)
			return
		}
		response.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	response.Write(w, http.StatusOK, models.ComputeResponse{
		Score:    result.Score,
		Category: result.Category,
		ZHR:      result.ZHR,
		ZRMSSD:   result.ZRMSSD,
	})
}
