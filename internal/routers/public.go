package routers

import (
	"ecm-api-template/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func PublicRouter(handler *handlers.Handlers) func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/", handler.HealthCheck)
		r.Route("/session", handler.Session.Route)
	}
}
