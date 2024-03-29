package middleware

import (
	"github.com/gin-gonic/gin"
	"stuLook-service/global"
	"stuLook-service/model"
	"stuLook-service/utils"

	"net/http"
)

func ApiAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取访问api的url和方法
		method, url := c.Request.Method, c.Request.URL.Path
		var api model.ApiMysql
		if err := global.ApiTable.Select("url,method").Where("url = ? and method = ?", url, method).Take(&api).Error; err != nil {
			c.JSON(http.StatusOK, utils.ErrorMess("验证api：此api不存在", err.Error()))
			c.Abort()
			return
		}
		//获取token解析出来的user
		userInterface, _ := c.Get("user")
		user := userInterface.(model.UserMysql)
		//获取user对应的role
		var role model.RoleMysql
		if err := global.RoleTable.Select("role").Where("role = ?", user.RoleId).Take(&role).Error; err != nil {
			//if err := global.RoleColl.FindOne(context.TODO(), bson.M{"_id": user.RoleId}).Decode(&role); err != nil {
			c.JSON(http.StatusOK, utils.ErrorMess("验证api：获取用户角色失败", err.Error()))
			c.Abort()
			return
		}
		//轮询role对应的apis，判断其是否相应的权限
		for i := range role.Api {
			if role.Api[i].Id == api.Id {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusOK, utils.ErrorMess("验证api：此用户无访问此api的权限", nil))
		c.Abort()
		return
	}
}
