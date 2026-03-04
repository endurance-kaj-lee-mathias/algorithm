package middleware

import (
	"log/slog"
	"net/http"

	"gitlab.com/kdg-ti/the-lab/teams-25-26/26-de-uitgeruste-it-ers/algorithm/internal/metrics/response"
)

func APIKeyAuth(apiKey string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			key := r.Header.Get("X-API-Key")
			if key == "" {
				response.WriteError(w, http.StatusUnauthorized, MissingAPIKey)
				slog.Warn("Missing API key in request")
				return
			}
			if key != apiKey {
				response.WriteError(w, http.StatusUnauthorized, InvalidAPIKey)
				slog.Warn("wrong API key in request")
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
