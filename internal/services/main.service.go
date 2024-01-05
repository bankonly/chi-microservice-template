package services

import (
	"ecm-api-template/internal/caches"
	"ecm-api-template/internal/repositories"
)

type Services struct {
	Session SessionService
}

func New(cache *caches.Caches, repo *repositories.Repositories) *Services {
	return &Services{
		Session: NewSessionService(cache.Session),
	}
}
