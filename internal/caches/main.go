package caches

import "github.com/redis/go-redis/v9"

type Caches struct {
	Session Session
}

func New(redisCli *redis.Client) *Caches {
	return &Caches{
		Session: NewSession(redisCli),
	}
}
