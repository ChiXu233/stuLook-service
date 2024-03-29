package utils

import (
	"github.com/gomodule/redigo/redis"
	"stuLook-service/global"
)

type Redis struct {
}

func (r Redis) SetValue(key, value string, t int) error {
	//6个小时和jwt一个时间
	if _, err := global.RedisPool.Get().Do("SET", key, value, "ex", t); err != nil {
		return err
	}
	return nil
}
func (r Redis) GetValue(key string) (string, error) {
	if value, err := redis.String(global.RedisPool.Get().Do("GET", key)); err != nil {
		return "", err
	} else {
		return value, nil
	}
}
