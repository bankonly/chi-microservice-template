package routers

import (
	"ecm-api-template/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func PublicRouter(r chi.Router) {
	handlers := handlers.NewHandlers()
	r.Get("/", handlers.HealthCheck)
	r.Route("/session", handlers.Session.Route)
}
