package initialize

import (
	"github.com/gomodule/redigo/redis"
	"stuLook-service/global"

	"log"
)

func RedisInit() {
	if global.RedisPool == nil {
		global.RedisPool = &redis.Pool{
			MaxIdle:     3, //最大空闲连接数
			MaxActive:   0, //最大链接数，0无限制
			IdleTimeout: 0, //连接不关闭
			Dial: func() (redis.Conn, error) {
				dial, err := redis.Dial("tcp", "43.138.43.184:6379", redis.DialPassword("root"))
				if err != nil {
					log.Println("redis连接失败", err.Error())
				}
				return dial, err
			},
		}
		global.RedisPool.Get()
	}
}
