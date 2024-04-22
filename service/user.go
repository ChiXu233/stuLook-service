package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strconv"
	"stuLook-service/global"
	"stuLook-service/model"
	"stuLook-service/utils"
	"time"
)

func CreateUser(user model.User) utils.Response {
	if err := global.UserColl.FindOne(context.TODO(), bson.M{"name": user.Name}).Decode(&bson.M{}); err == mongo.ErrNoDocuments {
		rand.Seed(time.Now().Unix()) //根据时间戳生成种子
		salt := strconv.FormatInt(rand.Int63(), 10)
		encryptedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password+salt), bcrypt.DefaultCost)
		if err != nil {
			return utils.ErrorMess("密码加密失败", err)
		}
		user.Password = string(encryptedPass)
		user.Salt = salt
		return utils.InsertOne(global.UserColl, user)
	} else {
		return utils.ErrorMess("创建失败", err.Error())
	}
}
func DeleteUser(_id primitive.ObjectID) utils.Response {
	return utils.DeleteOne(global.UserColl, bson.M{"_id": _id})
}
func GetUser(conditions map[string]interface{}, pageSize, currPage int64) utils.Response {
	return utils.GetPageData(global.UserColl, conditions, pageSize, currPage)
}
func Login(user model.User) utils.Response {
	var DBUser model.User
	//校验账号
	if err := global.UserColl.FindOne(context.TODO(), bson.M{"name": user.Name}).Decode(&DBUser); err != nil {
		return utils.ErrorMess("账号错误", err.Error())
	}
	//校验密码
	if err := bcrypt.CompareHashAndPassword([]byte(DBUser.Password), []byte(user.Password+DBUser.Salt)); err != nil {
		return utils.ErrorMess("密码错误", err.Error())
	}
	//生成token
	//token, err := ""
	////middleware.CreateToken(DBUser)
	//if err != nil {
	//	return utils.ErrorMess("生成token失败", err.Error())
	//}
	//res := map[string]interface{}{
	//	"_id":      DBUser.Id,
	//	"Name":     DBUser.Name,
	//	"password": DBUser.Password,
	//	//"token":    token,
	//}
	return utils.SuccessMess("登陆成功", "res")
}
