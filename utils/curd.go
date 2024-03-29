package utils

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 增加
func InsertOne(collection *mongo.Collection, document interface{}, opts ...*options.InsertOneOptions) Response {
	if res, err := collection.InsertOne(context.TODO(), document, opts...); err != nil {
		return ErrorMess("添加失败", err.Error())
	} else {
		return SuccessMess("添加成功", res)
	}
}

// 更新
func UpdateOne(collection *mongo.Collection, filter interface{}, update interface{}, opts ...*options.FindOneAndUpdateOptions) Response {
	if err := collection.FindOneAndUpdate(context.TODO(), filter, update, opts...).Decode(&bson.M{}); err != nil {
		return ErrorMess("更新失败", err.Error())
	} else {
		return SuccessMess("更新成功", filter)
	}
}

// 删除
func DeleteOne(collection *mongo.Collection, filter interface{}, opts ...*options.DeleteOptions) Response {
	res, err := collection.DeleteOne(context.TODO(), filter, opts...)
	if err != nil {
		return ErrorMess("删除失败", err.Error())
	}
	if res.DeletedCount == 0 {
		return ErrorMess("删除不存在", res)
	}
	return SuccessMess("删除成功", res)
}

// Find 查询
func Find(collection *mongo.Collection, result interface{}, filter interface{}, opts ...*options.FindOptions) error {
	cur, err := collection.Find(context.TODO(), filter, opts...)
	if err != nil {
		return err
	}
	if err = cur.All(context.TODO(), result); err != nil {
		return err
	}
	return nil
}

func GetPageData(collection *mongo.Collection, conditions map[string]interface{}, pageSize, currPage int64) Response {
	var pageData []map[string]interface{}
	skip := (currPage - 1) * pageSize
	//获取分页数据
	if err := Find(collection, &pageData, conditions, &options.FindOptions{
		Limit: &pageSize,
		Skip:  &skip,
		Sort:  bson.M{"_id": -1},
	}); err != nil {
		return ErrorMess("获取分页数据失败", err.Error())
	}
	//查询总数
	if total, err := collection.CountDocuments(context.TODO(), conditions); err != nil {
		return ErrorMess("获取总数时失败", err.Error())
	} else {
		res := map[string]interface{}{
			"pageData": pageData,
			"total":    total,
		}
		return SuccessMess("获取成功", res)
	}
}
