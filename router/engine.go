package router

import (
	"github.com/gin-gonic/gin"
	"stuLook-service/controller"
	"stuLook-service/middleware"
)

func GetEngine() *gin.Engine {
	engine := gin.Default()
	//路由分组
	engine.Use(middleware.Cors())           //跨域
	engine.POST("/login", controller.Login) //登录
	//engine.Use(middleware.JWTAuth(), middleware.ApiAuth()) //权限
	//Book
	BookRouter(engine)
	//user
	UserRouter(engine)
	//Api
	ApiRouter(engine)
	//Role
	RoleRouter(engine)
	return engine
}
