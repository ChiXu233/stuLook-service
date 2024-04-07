package initialize

func Init() {
	//MongoInit()
	RedisInit()
	MysqlInit()
	SnowFlakeInit()
	LogInit()
}
