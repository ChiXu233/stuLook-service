package initialize

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"stuLook-service/global"
)

func MongoInit() {
	if global.MongoClient == nil {
		//global.MongoClient = getMongoClient("mongodb://root:chixu%403048863634@43.138.43.184:27017")
		global.MongoClient = getMongoClient("mongodb://127.0.0.1:27017")
	}
	graphiteManager := global.MongoClient.Database("Book_library")
	{
		global.UserColl = graphiteManager.Collection("user")
		global.BookColl = graphiteManager.Collection("book")
		global.ApiColl = graphiteManager.Collection("api")
		global.RoleColl = graphiteManager.Collection("role")
	}
}

func getMongoClient(uri string) *mongo.Client {
	clientOptions := options.Client().ApplyURI(uri)
	MongoClient, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println(err)
	}
	if err = MongoClient.Ping(context.TODO(), nil); err != nil {
		log.Println(err)
	}
	return MongoClient
}
