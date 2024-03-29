package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stuLook-service/model"
	"stuLook-service/service"
	"stuLook-service/utils"
)

func CreateUserMysql(c *gin.Context) {
	var user model.UserMysql
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusOK, utils.ErrorMess("参数错误", err.Error()))
		return
	}
	c.JSON(http.StatusOK, service.CreateUserMysql(user))
}

func DeleteUserMysql(c *gin.Context) {
	_id := c.Query("id")
	if _id == "" {
		c.JSON(http.StatusOK, utils.ErrorMess("传入id格式错误", _id))
	}
	c.JSON(http.StatusOK, service.DeleteUserMysql(_id))

}

func GetUserMysql(c *gin.Context) {
	// 查询所有图书
	conditions := make(map[string]interface{})
	if name := c.Query("name"); name != "" {
		//i忽略大小写
		conditions["name"] = "%" + name + "%"
	} else {
		conditions["name"] = "%" + "" + "%"
	}
	pageSize := c.DefaultQuery("pageSize", "0")
	currPage := c.DefaultQuery("currPage", "1")
	c.JSON(http.StatusOK, service.GetUserMysql(conditions, pageSize, currPage))
}

func UpdateUserMysql(c *gin.Context) {
	var user model.UserMysql
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusOK, utils.ErrorMess("传入参数错误", err.Error()))
	}
	c.JSON(http.StatusOK, service.UpdateUserMysql(user))
}

//func LoginMysql(c *gin.Context) {
//	var user model.User
//	if err := c.Bind(&user); err != nil {
//		c.JSON(http.StatusOK, utils.ErrorMess("参数传递错误", err.Error()))
//	}
//	c.JSON(http.StatusOK, service.LoginMysql(user))
//}
