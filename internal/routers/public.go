package routers

import (
	"ecm-api-template/internal/caches"
	"ecm-api-template/internal/handlers"
	"ecm-api-template/internal/services"
	"ecm-api-template/pkg/storages"

	"github.com/go-chi/chi/v5"
)

func PublicRouter(r chi.Router) {
	cache := caches.New(storages.GetRedis())
	service := services.New(cache)
	handlers := handlers.New(service)

	r.Get("/", handlers.HealthCheck)
	r.Route("/session", handlers.Session.Route)
}
