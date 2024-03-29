package router

import (
	"github.com/gin-gonic/gin"
	"stuLook-service/controller"
)

func UserRouter(engine *gin.Engine) {
	user := engine.Group("user")
	{
		user.POST("create", controller.CreateUser)
		user.GET("get", controller.GetUser)
		user.DELETE("delete", controller.DeleteUser)

		user.POST("MysqlCreate", controller.CreateUserMysql)
		user.GET("MysqlGet", controller.GetUserMysql)
		user.DELETE("MysqlDelete", controller.DeleteUserMysql)
		user.PUT("MysqlUpdate", controller.DeleteUserMysql)
	}
}
