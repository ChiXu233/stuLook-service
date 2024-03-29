package router

import (
	"github.com/gin-gonic/gin"
	"stuLook-service/controller"
)

func RoleRouter(engine *gin.Engine) {
	role := engine.Group("role")
	{
		role.POST("create", controller.CreateRole)
		role.DELETE("delete", controller.DeleteRole)
		role.PUT("update", controller.UpdateRole)
		role.GET("get", controller.GetRole)

		role.POST("MysqlCreate", controller.CreateMysqlRole)
		role.DELETE("MysqlDelete", controller.DeleteMysqlRole)
		role.PUT("MysqlUpdate", controller.UpdateMysqlRole)
		role.GET("MysqlGet", controller.GetMysqlRole)
	}
}
