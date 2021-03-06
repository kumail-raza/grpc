package main

import (
	"fmt"
	"log"
	"net"

	"github.com/minhajuddinkhan/grpc/customer"
	"github.com/minhajuddinkhan/grpc/db"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	dbUserName          = "customerUser"
	dbPassword          = "customerPwd"
	dbName              = "grpc"
	customerServicePort = ":3443"
)

func main() {
	connectionString := fmt.Sprintf("mongodb://%s:%s@localhost:27017", dbUserName, dbPassword)
	mongoDB, err := db.NewMongoDB(connectionString, dbName)
	if err != nil {

		panic(err.Error())
	}

	s := grpc.NewServer()
	customer.RegisterCustomerServer(s, &customer.Server{Database: mongoDB})
	lis, err := net.Listen("tcp", customerServicePort)
	if err != nil {
		panic(err)
	}

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
