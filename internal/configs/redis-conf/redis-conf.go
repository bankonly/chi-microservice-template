package redisConf

import "time"

type RedisKeyConf struct {
	Key       string
	ExpireMin time.Duration
}

func Session(sessionId string) *RedisKeyConf {
	return &RedisKeyConf{
		Key:       "session:key_pair" + sessionId,
		ExpireMin: time.Minute * 120,
	}
}
