package controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strconv"
	"stuLook-service/model"
	"stuLook-service/service"
	"stuLook-service/utils"
)

func CreateUser(c *gin.Context) {
	var user model.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusOK, utils.ErrorMess("参数错误", err.Error()))
		return
	}
	c.JSON(http.StatusOK, service.CreateUser(user))
}

func DeleteUser(c *gin.Context) {
	if _id, err := primitive.ObjectIDFromHex(c.Query("_id")); err != nil {
		c.JSON(http.StatusOK, utils.ErrorMess("参数传递错误", err.Error()))
		return
	} else {
		c.JSON(http.StatusOK, service.DeleteUser(_id))
	}
}

func GetUser(c *gin.Context) {
	// 查询所有图书
	conditions := make(map[string]interface{})
	if name := c.Query("name"); name != "" {
		conditions["name"] = name
	}
	pageSize, err := strconv.ParseInt(c.DefaultQuery("pageSize", "0"), 10, 64)
	//获取 id 参数, 跟 Query 函数的区别是，可以通过第二个参数设置默认值。
	if err != nil {
		c.JSON(http.StatusOK, utils.ErrorMess(err.Error(), nil))
		return
	}
	currPage, err := strconv.ParseInt(c.DefaultQuery("currPage", "1"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, utils.ErrorMess(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, service.GetUser(conditions, pageSize, currPage))
}

func Login(c *gin.Context) {
	var user model.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusOK, utils.ErrorMess("参数传递错误", err.Error()))
	}
	c.JSON(http.StatusOK, service.Login(user))
}
