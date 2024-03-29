package controller

import (
	"github.com/gin-gonic/gin"
	"stuLook-service/model"
	"stuLook-service/service"
	"stuLook-service/utils"

	"net/http"
)

func CreateApiMysql(c *gin.Context) {
	var api model.ApiMysql
	if err := c.Bind(&api); err != nil {
		c.JSON(http.StatusOK, utils.ErrorMess("参数错误", err.Error()))
		return
	}
	c.JSON(http.StatusOK, service.CreateApiMysql(api))
}

func GetApiMysql(c *gin.Context) {
	conditions := make(map[string]interface{})
	if method := c.Query("method"); method != "" {
		conditions["method"] = method
	}
	if name := c.Query("name"); name != "" {
		//i忽略大小写
		conditions["name"] = "%" + name + "%"
	} else {
		conditions["name"] = "%" + "" + "%"
	}
	//默认获取全部数据
	pageSize := c.DefaultQuery("pageSize", "0")

	currPage := c.DefaultQuery("currPage", "1")

	c.JSON(http.StatusOK, service.GetApiMysql(pageSize, currPage, conditions))
}

func UpdateApiMysql(c *gin.Context) {
	var api model.ApiMysql
	if err := c.Bind(&api); err != nil {
		c.JSON(http.StatusOK, utils.ErrorMess("参数错误", err.Error()))
	}
	if api.Id == 0 {
		c.JSON(http.StatusOK, utils.ErrorMess("传入Id为空", nil))
	}
	c.JSON(http.StatusOK, service.UpdateApiMysql(api))
}

func DeleteApiMysql(c *gin.Context) {
	_id := c.Query("id")
	if _id == "" {
		c.JSON(http.StatusOK, utils.ErrorMess("传入id格式错误", _id))
	}
	c.JSON(http.StatusOK, service.DeleteApiMysql(_id))
}
