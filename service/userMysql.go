package service

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"stuLook-service/global"
	"stuLook-service/model"
	"stuLook-service/utils"

	"math/rand"
	"strconv"
	"time"
)

func CreateUserMysql(user model.UserMysql) utils.Response {
	if err := global.UserTable.Transaction(func(tx *gorm.DB) error {
		var userDB model.UserMysql
		if err := tx.Debug().Select("account").Where("account = ?", user.Account).Find(&userDB).Error; (err != nil && errors.Is(err, gorm.ErrRecordNotFound)) || userDB.Account == user.Account {
			return fmt.Errorf("账号重复:%w", err)
		}
		//查询角色是否存在
		var roleDB model.RoleMysql
		if err := global.RoleTable.Debug().Select("id").Where("id = ?", user.RoleId).Find(&roleDB).Error; err != nil {
			return fmt.Errorf("选择角色不存在:%w", err)
		}
		rand.New(rand.NewSource(time.Now().Unix()))
		//生成盐
		salt := strconv.FormatInt(rand.Int63(), 10)
		encryptedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password+salt), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("密码加密失败:%w", err)
		}
		user.Password, user.Salt = string(encryptedPass), salt
		//插入
		if err := tx.Create(&user).Error; err != nil {
			return fmt.Errorf("密码加密失败:%w", err)
		}
		return nil
	}); err != nil {
		return utils.ErrorMess("事务失败", err.Error())
	}
	return utils.SuccessMess("创建成功", nil)
}

func DeleteUserMysql(_id string) utils.Response {
	id, _ := strconv.ParseInt(_id, 10, 64)
	if err := global.UserTable.Transaction(func(tx *gorm.DB) error {
		if err := tx.Debug().Delete(&model.UserMysql{}, id).Error; err != nil {
			return fmt.Errorf("删除事务失败:%w", err)
		}
		return nil
	}); err != nil {
		return utils.ErrorMess("事务失败", err.Error())
	}
	return utils.SuccessMess("删除成功", id)
}

func GetUserMysql(conditions map[string]interface{}, pageSize, currPage string) utils.Response {
	skip, limit, err := utils.GetPage(currPage, pageSize)
	var count int64
	var dataDB []model.UserMysql
	if err != nil || limit == 0 {
		return utils.ErrorMess("传入分页失败", err.Error())
	}
	if err := global.UserTable.Transaction(func(tx *gorm.DB) error {
		res := tx.Debug().Order("id desc").Where("name LIKE ?", conditions["name"]).Preload("Role").Limit(limit).Offset(skip).Find(&dataDB).Count(&count)
		if res.Error != nil {
			return fmt.Errorf("查询事务执行失败:%w", res.Error)
		}
		return nil
	}); err != nil {
		return utils.ErrorMess("查询事务失败", err.Error())
	}
	return utils.SuccessMess("查询成功", struct {
		Count int64             `json:"count" bson:"count"`
		Data  []model.UserMysql `json:"data" bson:"data"`
	}{
		Count: count,
		Data:  dataDB,
	})
}

func UpdateUserMysql(user model.UserMysql) utils.Response {
	if err := global.UserTable.Transaction(func(tx *gorm.DB) error {
		var userDB model.UserMysql
		if err := tx.Where("id = ?", user.Id).First(&userDB).Error; err != nil {
			if err != nil {
				return fmt.Errorf("更新事务执行失败:%w", err)
			}
		}
		err := tx.Debug().Save(&user).Error
		if err != nil {
			return fmt.Errorf("更新用户失败:%w", err)
		}
		return nil
	}); err != nil {
		return utils.ErrorMess("更新业务失败", err.Error())
	}
	return utils.SuccessMess("更新成功", user.Id)
}

//func LoginMysql() utils.Response {
//
//}
