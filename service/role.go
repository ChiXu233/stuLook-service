package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"stuLook-service/global"
	"stuLook-service/model"
	"stuLook-service/utils"
)

func CreateRole(role model.Role) utils.Response {
	if err := global.RoleColl.FindOne(context.TODO(), bson.M{"name": role.Name}).Decode(bson.M{}); err != nil {
		if err == mongo.ErrNoDocuments {
			return utils.InsertOne(global.RoleColl, role)
		} else {
			return utils.ErrorMess("角色查询失败", err.Error())
		}
	} else {
		return utils.ErrorMess("角色重复", nil)
	}
}
func DeleteRole(_id primitive.ObjectID) utils.Response {
	return utils.DeleteOne(global.RoleColl, bson.M{"_id": _id})
}
func UpdateRole(role model.Role) utils.Response {
	for _, api := range role.Apis {
		if err := global.ApiColl.FindOne(context.TODO(), bson.M{"_id": api}).Decode(&bson.M{}); err != nil {
			return utils.ErrorMess(err.Error()+"此api不存在", api)
		}
	}
	//返回更新信息
	after := options.After
	err := global.RoleColl.FindOneAndUpdate(context.TODO(), bson.M{"_id": role.Id}, bson.M{"$set": role},
		&options.FindOneAndUpdateOptions{ReturnDocument: &after}).Decode(&role)
	if err != nil {
		return utils.ErrorMess("更新失败", err.Error())
	}
	return utils.SuccessMess("更新成功", err.Error())
}
func GetRole(conditions map[string]interface{}, pageSize, currPage int64) utils.Response {
	return utils.GetPageData(global.RoleColl, conditions, pageSize, currPage)
}
