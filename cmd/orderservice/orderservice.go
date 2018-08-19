package main

import (
	"context"
	"fmt"
	"log"

	"github.com/minhajuddinkhan/grpc-test/db"
	"github.com/mongodb/mongo-go-driver/bson"
)

func main() {

	userName := "ordersUser"
	password := "ordersPwd"
	dbName := "grpc"

	connectionString := fmt.Sprintf("mongodb://%s:%s@localhost:27017/grpc", userName, password)
	mongoDbConnection, err := db.NewMongoDB(connectionString)
	if err != nil {
		panic("Cannot connect to database")
	}

	var filter bson.Document
	err = mongoDbConnection.Client.Connect(context.Background())
	if err != nil {
		panic("Cannot connect to mongodb client" + err.Error())
	}

	cursor, err := mongoDbConnection.Client.Database(dbName).Collection("orders").Find(context.Background(), &filter)
	if err != nil {
		fmt.Println("error occured", err.Error())
	}
	for cursor.Next(context.Background()) {
		elem := bson.NewDocument()
		err := cursor.Decode(elem)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(elem)

	}

}
