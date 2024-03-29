package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Role struct {
	Id   primitive.ObjectID   `bson:"_id,omitempty" json:"_id,omitempty"`
	Name string               `bson:"name" json:"name"`
	Code string               `bson:"code" json:"code"`
	Apis []primitive.ObjectID `bson:"apis,omitempty" json:"apis,omitempty"`
	Desc string               `bson:"desc" json:"desc"`
}
