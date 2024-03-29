package initialize

func Init() {
	//initialize.MongoInit()
	//initialize.RedisInit()
	MysqlInit()
	SnowFlakeInit()
	LogInit()
}
