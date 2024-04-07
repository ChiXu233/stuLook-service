package utils

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"stuLook-service/global"
	"time"
)

var ctx = context.Background()

type Redis struct {
}

//func (r Redis) SetValue(key, value string, t int) error {
//	//6个小时和jwt一个时间
//	if _, err := global.RedisPool.Get().Do("SET", key, value, "ex", t); err != nil {
//		return err
//	}
//	return nil
//}
//func (r Redis) GetValue(key string) (string, error) {
//	if value, err := redis.String(global.RedisPool.Get().Do("GET", key)); err != nil {
//		return "", err
//	} else {
//		return value, nil
//	}
//}

// GetRedisType Type用来获取一个key对应值的类型
func (r Redis) GetRedisType(key string) string {
	vType, err := global.RedisClient.Type(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	return vType
}

// DelRedisKey 删除缓存项
func (r Redis) DelRedisKey(keys []string) {
	n, err := global.RedisClient.Del(ctx, keys...).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("成功删除了 %v 个\n", n)
}

// ExistsRedisKey 检测key是否存在
func (r Redis) ExistsRedisKey(key string) bool {
	//注：Exists()方法可以传入多个key,返回的第一个结果表示存在的key的数量,不过工作中我们一般不同时判断多个key是否存在，一般就判断一个key,所以判断是否大于0即可，如果判断判断传入的多个key是否都存在，则返回的结果的值要和传入的key的数量相等
	n, err := global.RedisClient.Exists(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	//若n>0则存在
	return n > 0
}

// ExpireRedisKey Expire()方法是设置某个时间段(time.Duration)后过期
func (r Redis) ExpireRedisKey(key string, t time.Duration) {
	res, err := global.RedisClient.Expire(ctx, key, t).Result()
	if err != nil {
		panic(err)
	}
	if res {
		fmt.Println("设置成功")
	} else {
		fmt.Println("设置失败")
	}
}

// ExpireAtRedisKey ExpireAt()方法是在某个时间点(time.Time)过期失效
func (r Redis) ExpireAtRedisKey(key string, t time.Time) {
	res, err := global.RedisClient.ExpireAt(ctx, key, t).Result()
	if err != nil {
		panic(err)
	}
	if res {
		fmt.Println("设置成功")
	} else {
		fmt.Println("设置失败")
	}
}

// TTLRedisKey TTL()方法可以获取某个键的剩余有效期单位：秒； PTTL()获取毫秒
func (r Redis) TTLRedisKey(key string) time.Duration {
	ttl, err := global.RedisClient.TTL(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	return ttl
}

// RedisDBSize DBSize()查看当前数据库key的数量
func (r Redis) RedisDBSize() int64 {
	num, err := global.RedisClient.DBSize(ctx).Result()
	if err != nil {
		panic(err)
	}
	return num
}

// RedisFlushDB FlushDB():清空当前数据库
func (r Redis) RedisFlushDB() string {
	///清空当前数据库，因为连接的是索引为0的数据库，所以清空的就是0号数据库
	res, err := global.RedisClient.FlushDB(ctx).Result()
	if err != nil {
		panic(err)
	}
	return res //OK
}

// RedisSetValue Set():设置;仅仅支持字符串(包含数字)操作，不支持内置数据编码功能。如果需要存储Go的非字符串类型，需要提前手动序列化，获取时再反序列化。
func (r Redis) RedisSetValue(key, value string, t time.Duration) {
	err := global.RedisClient.Set(ctx, key, value, t).Err()
	if err != nil {
		panic(err)
	}
}

// RedisSetValueEx SetEX():设置并指定过期时间
func (r Redis) RedisSetValueEx(key, value string, t time.Duration) {
	err := global.RedisClient.SetEx(ctx, key, value, t).Err()
	if err != nil {
		panic(err)
	}
}

// RedisGetValue Get():获取
func (r Redis) RedisGetValue(key string) string {
	value, err := global.RedisClient.Get(ctx, key).Result()
	if err != redis.Nil && err != nil {
		fmt.Println("key不存在")
		panic(err)
	}
	return value
}
