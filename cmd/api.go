package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gitlab.com/kdg-ti/the-lab/teams-25-26/26-de-uitgeruste-it-ers/algorithm/cmd/auth"
	"gitlab.com/kdg-ti/the-lab/teams-25-26/26-de-uitgeruste-it-ers/algorithm/internal/health"
	message "gitlab.com/kdg-ti/the-lab/teams-25-26/26-de-uitgeruste-it-ers/algorithm/internal/metrics"
)

func (server *server) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(time.Minute))

	handler := message.Wire()
	healthHandler := health.NewHandler()

	r.Use(auth.TokenAuthentication(server.idp))

	r.Post("/metrics", handler.CreateMessage)
	r.Get("/health", healthHandler.Health)

	return r
}

func (server *server) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         server.config.Port,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	slog.Info("server has started", "port", server.config.Port)
	return srv.ListenAndServe()
}
