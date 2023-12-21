package services

import "ecm-api-template/internal/caches"

type Services struct {
	Session SessionService
}

func New(cache *caches.Caches) *Services {
	return &Services{
		Session: NewSessionService(cache.Session),
	}
}
