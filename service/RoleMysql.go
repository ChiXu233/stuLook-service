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

func CreateMysqlRole(role model.RoleMysql) utils.Response {
	if err := global.RoleTable.Transaction(func(tx *gorm.DB) error {
		var roleDB model.RoleMysql
		if err := tx.Debug().Select("name").Where("name = ?", role.Name).Find(&roleDB).Error; (err != nil && !errors.Is(err, gorm.ErrRecordNotFound)) || roleDB.Name == role.Name {
			return fmt.Errorf("创建角色事务失败:%w", err)
		}
		//查询role中api是否正确
		var apiDB []model.ApiMysql
		if err := global.ApiTable.Select("id").Where("id IN ?", extractRoleId(role.Api)).Find(&apiDB).Error; err != nil {
			return fmt.Errorf("查询api错误:%w", err)
		}
		if len(role.Api) != len(apiDB) {
			return fmt.Errorf("api数量不等")
		}
		role.Id = global.RoleSnowFlake.Generate().Int64()
		if err := tx.Debug().Create(&role).Error; err != nil {
			return fmt.Errorf("创建角色失败:%w", err)
		}
		return nil
	}); err != nil {
		return utils.ErrorMess("创建角色失败", err.Error())
	}
	return utils.SuccessMess("创建角色成功", nil)
}

func DeleteMysqlRole(id string) utils.Response {
	_id, _ := strconv.ParseInt(id, 10, 64)
	if err := global.RoleTable.Transaction(func(tx *gorm.DB) error {
		var roleDB model.RoleMysql
		if err := tx.Debug().Where("id = ?", _id).Find(&roleDB).Error; err != nil {
			return fmt.Errorf("传入角色id错误:%w", err)
		}
		tx0 := global.ApiTable.Begin()
		if err := tx0.Model(&model.RoleMysql{Id: _id}).Association("Api").Clear(); err != nil {
			tx0.Rollback()
			return fmt.Errorf("删除角色api关联失败:%w", err)
		}
		tx0.Commit()
		//删除角色
		if err := tx.Debug().Delete(&model.RoleMysql{}, _id).Error; err != nil {
			return fmt.Errorf("删除角色失败:%w", err)
		}
		return nil
	}); err != nil {
		return utils.ErrorMess("创建角色失败", err.Error())
	}
	return utils.SuccessMess("创建角色成功", nil)
}

func UpdateMysqlRole(role model.RoleMysql) utils.Response {
	if err := global.RoleTable.Transaction(func(tx *gorm.DB) error {
		var roleDB model.RoleMysql
		if err := tx.Debug().Where("id = ?", role.Id).Find(&roleDB).Error; err != nil {
			return fmt.Errorf("查询失败%w", err)
		}
		if err := tx.Save(&role).Error; err != nil {
			return fmt.Errorf("更新角色失败:%w", err)
		}
		return nil
	}); err != nil {
		return utils.ErrorMess("创建角色失败", err.Error())
	}
	return utils.SuccessMess("创建角色成功", nil)
}
func GetMysqlRole(conditions map[string]interface{}, currPage, pageSize, startTime, endTime string) utils.Response {
	skip, limit, err := utils.GetPage(currPage, pageSize)
	var count int64
	var dataDB []model.RoleMysql
	if err != nil || limit == 0 {
		return utils.ErrorMess("传入分页失败", err.Error())
	}
	if err := global.RoleTable.Transaction(func(tx *gorm.DB) error {
		res := tx.Debug().Order("id desc").Where("name LIKE ?", conditions["name"]).Preload("Api").Preload("User").Limit(limit).Offset(skip).Find(&dataDB).Count(&count)
		if res.Error != nil {
			return fmt.Errorf("查询事务执行失败:%w", res.Error)
		}
		return nil
	}); err != nil {
		return utils.ErrorMess("查询事务失败", err.Error())
	}
	return utils.SuccessMess("查询成功", struct {
		Count int64             `json:"count" bson:"count"`
		Data  []model.RoleMysql `json:"data" bson:"data"`
	}{
		Count: count,
		Data:  dataDB,
	})
}

// 提取role中的apiID
func extractRoleId(apis []model.ApiMysql) []int64 {
	ids := make([]int64, len(apis))
	for i, v := range apis {
		ids[i] = v.Id
	}
	return ids
}
