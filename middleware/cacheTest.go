package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"stuLook-service/utils"
	"sync"
	"time"
)

// 对gin框架http响应的二次封装
type bodyLogWriter struct {
	mutex sync.Mutex
	gin.ResponseWriter
	body *bytes.Buffer
}

// 将传入的字节切片写入到 bodyLogWriter 结构体中的缓冲区 body 中，并同时写入到原始的 gin.ResponseWriter 中，保证了响应的正常输出，并通过互斥锁确保了在并发环境下的安全性
func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	n, err := w.body.Write(b)
	if err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(b)
}

func CacheTest() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		path := c.Request.URL.Path
		cache := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = cache
		if method == "GET" && c.Writer.Status() == 200 {
			cacheData, err := utils.Redis{}.RedisGetValue(path)
			if err != nil {
				logrus.Error("获取redis缓存未命中:", err)
				c.Next()
				if cache.body != nil {
					errOver := utils.Redis{}.RedisSetValue(path, cache.body.String(), 5*time.Minute)
					if errOver != nil {
						logrus.Error("更新redis缓存失败", err)
					}
				}
				return
			}
			//若缓存命中
			if cacheData != "" {
				var jsonData interface{}
				err := json.Unmarshal([]byte(cacheData), &jsonData)
				if err != nil {
					logrus.Error("缓存解码失败", err)
					c.Next()
					return
				}
				c.JSON(http.StatusOK, jsonData)
				c.Abort()
				return
			}
		}
		//延迟双删
		err := utils.Redis{}.DelRedisKey(path)
		if err != nil {
			logrus.Error("删除redis缓存错误", err)
			return
		}
		c.Next()
		err = utils.Redis{}.DelRedisKey(path)
		if err != nil {
			logrus.Error("删除redis缓存错误", err)
			return
		}
	}
}
