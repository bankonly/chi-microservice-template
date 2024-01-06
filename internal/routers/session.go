package routers

import (
	"ecm-api-template/internal/handlers"
	"ecm-api-template/internal/middlewares"

	"github.com/go-chi/chi/v5"
)

func SessionRouter(handler *handlers.Handlers) func(r chi.Router) {
	return func(r chi.Router) {
		r.Use(middlewares.VerifySession)
		r.Route("/", handler.Session.Route)
	}
}
