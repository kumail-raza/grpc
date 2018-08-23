package main

import (
	"fmt"
	"log"
	"net"

	"github.com/minhajuddinkhan/grpc/db"
	"github.com/minhajuddinkhan/grpc/order"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	dbUserName       = "ordersUser"
	dbPassword       = "ordersPwd"
	dbName           = "grpc"
	orderServiceAddr = ":4443"
)

func main() {
	connectionString := fmt.Sprintf("mongodb://%s:%s@localhost:27017", dbUserName, dbPassword)
	mongoDB, err := db.NewMongoDB(connectionString, dbName)
	if err != nil {

		panic(err.Error())
	}

	s := grpc.NewServer()
	order.RegisterOrderServer(s, &order.Server{Database: mongoDB})
	lis, err := net.Listen("tcp", orderServiceAddr)
	if err != nil {
		panic(err)
	}

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
