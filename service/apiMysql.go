package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"stuLook-service/global"
	"stuLook-service/model"
	"stuLook-service/utils"

	"strconv"
)

func CreateApiMysql(api model.ApiMysql) utils.Response {
	if err := global.ApiTable.Transaction(func(tx *gorm.DB) error {
		//查询api是否重复
		var apiDB model.ApiMysql
		if err := tx.Select("url").Where("url = ?", api.Url).First(&apiDB).Error; (err != nil && !errors.Is(err, gorm.ErrRecordNotFound)) || api.Url == apiDB.Url {
			return fmt.Errorf("api重复:%w", err)
		}
		api.Id = global.ApiSnowFlake.Generate().Int64()
		if err := tx.Create(&api).Error; err != nil {
			return fmt.Errorf("api创建失败:%w", err)
		}
		return nil
	}); err != nil {
		return utils.ErrorMess("事务失败", err.Error())
	}
	return utils.SuccessMess("插入成功", api)
}

func GetApiMysql(pageSize, currPage string, conditions map[string]interface{}) utils.Response {
	skip, limit, err := utils.GetPage(currPage, pageSize)
	var count int64
	var dataDB []model.ApiMysql
	if err != nil || limit == 0 {
		return utils.ErrorMess("传入分页失败", err.Error())
	}
	if err := global.ApiTable.Transaction(func(tx *gorm.DB) error {
		res := tx.Debug().Order("id desc").Where("name LIKE ?", conditions["name"]).Or("method = ?", conditions["method"]).Limit(limit).Offset(skip).Find(&dataDB).Count(&count)
		if res.Error != nil {
			return fmt.Errorf("查询事务执行失败:%w", res.Error)
		}
		return nil
	}); err != nil {
		return utils.ErrorMess("查询事务失败", err.Error())
	}
	return utils.SuccessMess("查询成功", struct {
		Count int64            `json:"count" bson:"count"`
		Data  []model.ApiMysql `json:"data" bson:"data"`
	}{
		Count: count,
		Data:  dataDB,
	})
}

func UpdateApiMysql(api model.ApiMysql) utils.Response {
	if err := global.ApiTable.Transaction(func(tx *gorm.DB) error {
		var apiDB model.ApiMysql
		if err := tx.Debug().Where("id = ?", api.Id).First(&apiDB).Error; err != nil || errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("更新事务失败%w", err)
		}
		if err := tx.Save(&api).Error; err != nil {
			return fmt.Errorf("更新事务失败%w", err)
		}
		return nil
	}); err != nil {
		return utils.ErrorMess("更新失败", err.Error())
	}
	return utils.SuccessMess("更新api成功", api.Id)
}

func DeleteApiMysql(_id string) utils.Response {
	id, err := strconv.ParseInt(_id, 10, 64)
	if err != nil {
		return utils.ErrorMess("字符串转化整数失败", err.Error())
	}
	if err := global.ApiTable.Transaction(func(tx *gorm.DB) error {
		//删除对应角色中的api（新开启事务）
		tx0 := global.RoleTable.Begin()
		//清除角色表api相关联
		if err := tx0.Model(&model.ApiMysql{Id: id}).Association("RoleMysql").Clear(); err != nil {
			tx0.Rollback()
			return fmt.Errorf("清除关联失败:%w", err)
		}
		tx0.Commit()
		if err := tx.Debug().Delete(&model.ApiMysql{}, id).Error; err != nil || errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("删除事务失败%w", err)
		}
		return nil
	}); err != nil {
		return utils.ErrorMess("删除事务失败", err.Error())
	}
	return utils.SuccessMess("删除成功", _id)

}
