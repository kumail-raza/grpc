package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/minhajuddinkhan/grpc-test/customer"
	customerproto "github.com/minhajuddinkhan/grpc-test/customer/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/minhajuddinkhan/grpc-test/db"
)

func main() {

	const (
		dbUserName = "customerUser"
		dbPassword = "customerPwd"
		dbName     = "grpc"
	)

	connectionString := fmt.Sprintf("mongodb://%s:%s@localhost:27017", dbUserName, dbPassword)
	mongoDB, err := db.NewMongoDB(connectionString, dbName)
	if err != nil {

		panic(err.Error())
	}
	customerRouter := customer.NewCustomerRouter(mongoDB)

	s := grpc.NewServer()
	customerproto.RegisterCustomerServer(s, &customer.Server{Database: mongoDB})
	lis, err := net.Listen("tcp", ":3443")
	if err != nil {
		panic(err)
	}

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	http.ListenAndServe(":3000", customerRouter)

}
