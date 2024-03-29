package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Api struct {
	Id         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name       string             `bson:"name" json:"name"`     //api名称
	Url        string             `bson:"url" json:"url"`       //api路由
	Method     string             `bson:"method" json:"method"` //方法
	Desc       string             `bson:"desc" json:"desc"`     //描述
	UpdateTime string             `bson:"updateTime" json:"updateTime"`
	CreateTime string             `bson:"createTime" json:"createTime"`
}
