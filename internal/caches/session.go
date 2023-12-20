package caches

import (
	"context"
	"ecm-api-template/internal/configs"
	"time"

	"github.com/bankonly/go-pkg/v1/writers"
	"github.com/redis/go-redis/v9"
)

type Session interface {
	Save(requestId, sessionId, publicKey string) error
	GetSession(requestId, sessionId string) string
	SaveVector(requestId, vector string) error
	GetVector(requestId, vector string) string
}

type SessionOpts struct {
	redisCli *redis.Client
}

func NewSession(redisCli *redis.Client) Session {
	return &SessionOpts{redisCli: redisCli}
}

func (opts *SessionOpts) Save(requestId, sessionId, publicKey string) error {
	key := configs.RedisConf.SessionInfo.Key + sessionId
	err := opts.redisCli.Set(context.Background(), key, []byte(publicKey), time.Minute*configs.RedisConf.SessionInfo.ExpireMin).Err()
	if err != nil {
		writers.Console(requestId, "Cache.Session.Save: "+err.Error())
	}
	return err
}

func (opts *SessionOpts) GetSession(requestId, sessionId string) string {
	key := configs.RedisConf.SessionInfo.Key + sessionId
	result, err := opts.redisCli.Get(context.Background(), key).Result()
	if err != nil {
		writers.Console(requestId, "Cache.Session.GetSession: "+err.Error())
	}
	return result
}

func (opts *SessionOpts) SaveVector(requestId, vector string) error {
	key := configs.RedisConf.SessionVector.Key + vector
	err := opts.redisCli.Set(context.Background(), key, []byte(vector), time.Minute*configs.RedisConf.SessionVector.ExpireMin).Err()
	if err != nil {
		writers.Console(requestId, "Cache.Session.SaveVector: "+err.Error())
	}
	return err
}

func (opts *SessionOpts) GetVector(requestId, vector string) string {
	key := configs.RedisConf.SessionVector.Key + vector
	result, err := opts.redisCli.Get(context.Background(), key).Result()
	if err != nil {
		writers.Console(requestId, "Cache.Session.GetVector: "+err.Error())
	}
	return result
}
