package controller

import (
	"github.com/gin-gonic/gin"
	"stuLook-service/model"
	"stuLook-service/service"
	"stuLook-service/utils"

	"net/http"
)

func CreateMysqlRole(c *gin.Context) {
	var role model.RoleMysql
	if err := c.Bind(&role); err != nil {
		c.JSON(http.StatusOK, utils.ErrorMess("参数传递错误", err.Error()))
		return
	}
	c.JSON(http.StatusOK, service.CreateMysqlRole(role))
}

func DeleteMysqlRole(c *gin.Context) {
	id := c.Query("id")
	if id == " " {
		c.JSON(http.StatusOK, utils.ErrorMess("id参数传递错误", nil))
		return
	}
	c.JSON(http.StatusOK, service.DeleteMysqlRole(id))
}

func UpdateMysqlRole(c *gin.Context) {
	var role model.RoleMysql
	if err := c.Bind(&role); err != nil {
		c.JSON(http.StatusOK, utils.ErrorMess("参数传递错误", err.Error()))
		return
	}
	c.JSON(http.StatusOK, service.UpdateMysqlRole(role))
}

func GetMysqlRole(c *gin.Context) {
	conditions := make(map[string]interface{})
	if name := c.Query("name"); name != "" {
		//i忽略大小写
		conditions["name"] = "%" + name + "%"
	} else {
		conditions["name"] = "%" + "" + "%"
	}
	currPage := c.DefaultQuery("currPage", "1")
	pageSize := c.DefaultQuery("pageSize", "10")
	startTime := c.Query("startTime")
	endTime := c.Query("endTime")
	c.JSON(http.StatusOK, service.GetMysqlRole(conditions, currPage, pageSize, startTime, endTime))
}
