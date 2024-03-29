package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type ApiUpdate struct {
	Id         primitive.ObjectID `bson:"id,omitempty" json:"id,omitempty"`
	Name       string             `bson:"name" json:"name"`     //api名称
	Url        string             `bson:"url" json:"url"`       //api路由
	Method     string             `bson:"method" json:"method"` //方法
	Desc       string             `bson:"desc" json:"desc"`     //描述
	UpdateTime string             `bson:"updateTime" json:"updateTime"`
}
