package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	Id         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name       string             `bson:"name" json:"name"`
	Author     string             `bson:"author" json:"author"`
	Details    string             `bson:"details" json:"details"`
	UpdateTime string             `bson:"updateTime" json:"updateTime"`
	CreateTime string             `bson:"createTime" json:"createTime"`
}
