package initialize

import (
	"stuLook-service/global"
)

func Init() {
	//MongoInit()
	RedisInit()
	MysqlInit()
	SnowFlakeInit()
	LogInit()
	InitPool()
	//if runtime.GOOS != "linux" {
	//	return
	//}
	MqttInit("StuLook", &global.MqTest)
}
