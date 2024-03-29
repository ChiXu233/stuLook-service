package service

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"stuLook-service/global"
	"stuLook-service/model"
	"stuLook-service/utils"
	"time"
)

func CreateBook(book model.Book) utils.Response {
	book.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	return utils.InsertOne(global.BookColl, book)
}
func DeleteBook(_id primitive.ObjectID) utils.Response {
	return utils.DeleteOne(global.BookColl, bson.M{"_id": _id})
}
func UpdateBook(book model.BookUpdate) utils.Response {
	book.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
	return utils.UpdateOne(global.BookColl, bson.M{"_id": book.Id}, bson.M{"$set": book})
}
func GetBook(conditions map[string]interface{}, pageSize, currPage int64) utils.Response {
	return utils.GetPageData(global.BookColl, conditions, pageSize, currPage)
}
func GetBookById(_id primitive.ObjectID) utils.Response {
	var result bson.M
	if err := global.BookColl.FindOne(context.TODO(), bson.M{"_id": _id}).Decode(&result); err != nil {
		return utils.ErrorMess("查找失败", err.Error())
	}
	//if err != nil {
	//	fmt.Println(result)
	//	return utils.SuccessMess("查找成功", result)
	//}
	return utils.SuccessMess("查找成功", result)
}

//	func GetByName(name string) utils.Response {
//		var books []model.Book
//		if result, err := global.BookColl.Find(context.TODO(), bson.M{"name": bson.M{"$regex": primitive.Regex{Pattern: name, Options: "i"}}}); err != nil {
//			return utils.ErrorMess("查找失败", err.Error())
//		}
//		if err := c.All(context.Background(), &books); err != nil {
//			return nil, err
//		}
//		return utils.SuccessMess("查找成功", result)
//	}
func GetByName(name string) utils.Response {
	cursor, err := global.BookColl.Find(context.TODO(), bson.M{"name": bson.M{"$regex": primitive.Regex{Pattern: name, Options: "i"}}})
	if err != nil {
		return utils.ErrorMess("模糊查找失败", err.Error())
	}
	var books []model.Book
	if err := cursor.All(context.TODO(), &books); err != nil {
		return utils.ErrorMess("数组赋值失败", err.Error())
	}
	return utils.SuccessMess("模糊查询成功", books)
}

func GetByAuthor(Author string) utils.Response {
	fmt.Println("44444444")
	cursor, err := global.BookColl.Find(context.TODO(), bson.M{"author": bson.M{"$regex": primitive.Regex{Pattern: Author, Options: "i"}}})
	if err != nil {
		return utils.ErrorMess("模糊查找失败", err.Error())
	}
	var Authors []model.Book
	if err := cursor.All(context.TODO(), &Authors); err != nil {
		return utils.ErrorMess("数组赋值失败", err.Error())
	}
	return utils.SuccessMess("模糊查询成功", Authors)
}
