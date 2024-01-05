package routers

import (
	"ecm-api-template/internal/caches"
	"ecm-api-template/internal/handlers"
	"ecm-api-template/internal/repositories"
	"ecm-api-template/internal/services"
	"ecm-api-template/pkg/storages"

	"github.com/go-chi/chi/v5"
)

func Router(r chi.Router) {
	cache := caches.New(storages.GetRedis())
	repo := repositories.New(storages.GetPostgresDB())
	service := services.New(cache, repo)
	handlers := handlers.New(repo, service)

	r.Route("/", PublicRouter(handlers))
}
