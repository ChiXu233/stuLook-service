package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"stuLook-service/model"
	"stuLook-service/service"
	"stuLook-service/utils"

	"net/http"
	"strconv"
)

func CreateRole(c *gin.Context) {
	var role model.Role
	if err := c.Bind(&role); err != nil {
		c.JSON(http.StatusOK, utils.ErrorMess("参数传递错误", err.Error()))
	}
	c.JSON(http.StatusOK, service.CreateRole(role))
}
func DeleteRole(c *gin.Context) {
	if _id, err := primitive.ObjectIDFromHex(c.Query("_id")); err != nil {
		c.JSON(http.StatusOK, utils.ErrorMess("参数传递错误", err.Error()))
	} else {
		c.JSON(http.StatusOK, service.DeleteRole(_id))
	}
}
func UpdateRole(c *gin.Context) {
	var role model.Role
	var _id, err = primitive.ObjectIDFromHex(c.Query("_id"))
	fmt.Println(err)
	if err1 := c.ShouldBind(&role); err1 != nil {
		c.JSON(http.StatusOK, utils.ErrorMess("参数错误", err1.Error()))
		return
	}
	role.Id = _id
	c.JSON(http.StatusOK, service.UpdateRole(role))
}
func GetRole(c *gin.Context) {
	conditions := make(map[string]interface{})
	if method := c.Query("method"); method != "" {
		conditions["method"] = method
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
	c.JSON(http.StatusOK, service.GetRole(conditions, pageSize, currPage))
}
