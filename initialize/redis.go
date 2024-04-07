package initialize

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"stuLook-service/global"
)

// Background返回一个非空的Context。 它永远不会被取消，没有值，也没有期限。
// 它通常在main函数，初始化和测试时使用，并用作传入请求的顶级上下文。
var ctx = context.Background()

//func RedisInit() {
//	if global.RedisPool == nil {
//		global.RedisPool = &redis.Pool{
//			MaxIdle:     3, //最大空闲连接数
//			MaxActive:   0, //最大链接数，0无限制
//			IdleTimeout: 0, //连接不关闭
//			Dial: func() (redis.Conn, error) {
//				dial, err := redis.Dial("tcp", "43.138.43.184:6379", redis.DialPassword("root"))
//				if err != nil {
//					log.Println("redis连接失败", err.Error())
//				}
//				return dial, err
//			},
//		}
//		global.RedisPool.Get()
//	}
//}

func RedisInit() {
	global.RedisClient = redis.NewClient(&redis.Options{
		Addr:         "43.138.43.184:6379", // Redis 服务器地址
		Password:     "root",               // Redis 服务器密码
		DB:           0,                    // Redis 数据库索引
		PoolSize:     5,                    // 连接池大小
		MinIdleConns: 3,                    // 最小空闲连接数
	})

	//defer global.RedisClient.Close()
	pong, err := global.RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Println("redis连接失败", pong, err)
	}
	fmt.Println("redis连接成功", pong)
}
