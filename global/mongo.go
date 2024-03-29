package global

import "go.mongodb.org/mongo-driver/mongo"

var (
	MongoClient *mongo.Client
	BookColl    *mongo.Collection
	UserColl    *mongo.Collection
	ApiColl     *mongo.Collection
	RoleColl    *mongo.Collection
)
