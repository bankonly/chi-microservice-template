package services

import "ecm-api-template/internal/caches"

type Services struct {
	Session Session
}

func New(cache *caches.Caches) *Services {
	return &Services{
		Session: NewSession(cache.Session),
	}
}
