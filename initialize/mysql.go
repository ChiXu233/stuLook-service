package initialize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"stuLook-service/global"

	"log"
	"time"
)

const dsn = "root:123123@tcp(localhost:3306)/stulook?charset=utf8&parseTime=True&loc=Local"

func MysqlInit() {
	var err error
	global.MysqlClient, err = gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 171,
	}), &gorm.Config{
		SkipDefaultTransaction: false,
		//NamingStrategy: schema.NamingStrategy{
		//	//TablePrefix:   "StuLook_", // table name prefix, table for `User` would be `t_users`
		//	SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		//},
		DisableForeignKeyConstraintWhenMigrating: true, //逻辑外键（代码里自动体现外键关系）
	})
	if err != nil {
		log.Fatalln("Mysql数据库连接失败:", err)
	}
	sqlDB, err := global.MysqlClient.DB()
	if err != nil {
		log.Fatalln("连接池创建失败")
	}
	sqlDB.SetMaxIdleConns(10)                  //最大空闲连接数
	sqlDB.SetMaxOpenConns(10)                  //最大连接数
	sqlDB.SetConnMaxLifetime(time.Minute * 15) //设置连接空闲超时
	{
		global.UserTable = global.MysqlClient.Table("user_mysqls")
		global.RoleTable = global.MysqlClient.Table("role_mysqls")
		global.ApiTable = global.MysqlClient.Table("api_mysqls")
		global.LogTable = global.MysqlClient.Table("log")
	}
	fmt.Println("mysql连接成功")
}
