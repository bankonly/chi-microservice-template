package caches

import (
	redisConf "ecm-api-template/internal/configs/redis-conf"

	"github.com/bankonly/go-pkg/v1/writers"
	"github.com/go-redis/redis"
)

type Session interface {
	Save(requestId, sessionId, publicKey string) error
	GetSession(requestId, sessionId string) string
}

type SessionOpts struct {
	redisCli *redis.Client
}

func NewSession() Session {
	return &SessionOpts{}
}

func (opts *SessionOpts) Save(requestId, sessionId, publicKey string) error {
	conf := redisConf.Session(sessionId)
	err := opts.redisCli.Set(conf.Key, []byte(publicKey), conf.ExpireMin).Err()
	if err != nil {
		writers.Console(requestId, "Cache.Session.Save: "+err.Error())
	}
	return err
}

func (opts *SessionOpts) GetSession(requestId, sessionId string) string {
	conf := redisConf.Session(sessionId)
	result, err := opts.redisCli.Get(conf.Key).Result()
	if err != nil {
		writers.Console(requestId, "Cache.Session.GetSession: "+err.Error())
	}
	return result
}
