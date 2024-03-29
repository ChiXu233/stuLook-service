package router

import (
	"github.com/gin-gonic/gin"
	"stuLook-service/controller"
)

func BookRouter(engine *gin.Engine) {
	book := engine.Group("book")
	{
		book.POST("/create", controller.CreateBook)      //增加
		book.GET("/get", controller.GetBook)             //查询所有
		book.PUT("/update", controller.UpdateBook)       //更新
		book.DELETE("/delete", controller.DeleteBook)    //删除
		book.GET("/getById", controller.GetById)         //根据ID查找
		book.GET("/getByName", controller.GetByName)     //根据图书名模糊查询
		book.GET("/getByAuthor", controller.GetByAuthor) //根据作者模糊查询
	}
}
