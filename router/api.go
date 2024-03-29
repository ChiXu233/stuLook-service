package router

import (
	"github.com/gin-gonic/gin"
	"stuLook-service/controller"
)

func ApiRouter(engine *gin.Engine) {
	api := engine.Group("api")
	{
		api.GET("get", controller.GetApi)
		api.POST("create", controller.CreateApi)
		api.DELETE("delete", controller.DeleteApi)
		api.PUT("update", controller.UpdateApi)

		api.POST("/MysqlCreate", controller.CreateApiMysql)
		api.GET("/MysqlGet", controller.GetApiMysql)
		api.PUT("/MysqlUpdate", controller.UpdateApiMysql)
		api.DELETE("/MysqlDelete", controller.DeleteApiMysql)
	}
}
