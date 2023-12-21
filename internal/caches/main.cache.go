package caches

import "github.com/redis/go-redis/v9"

type Caches struct {
	Session SessionCache
}

func New(redisCli *redis.Client) *Caches {
	return &Caches{
		Session: NewSessionCache(redisCli),
	}
}
