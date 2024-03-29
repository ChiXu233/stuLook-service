package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name     string             `bson:"name" json:"name"`
	Account  string             `bson:"account" json:"account"`
	Password string             `bson:"password,omitempty" json:"password,omitempty"`
	Salt     string             `bson:"salt,omitempty" json:"salt,omitempty"`
	RoleId   primitive.ObjectID `bson:"roleId" json:"roleId"`
}
