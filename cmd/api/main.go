package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/minhajuddinkhan/grpc/protocols"
	"github.com/minhajuddinkhan/grpc/routes"
	"google.golang.org/grpc"
)

const (
	apiGatewayPort           = ":3000"
	customerMicroServiceAddr = "localhost:3443"
)

func main() {

	custConn, err := grpc.Dial(customerMicroServiceAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer custConn.Close()
	c := customerprotocol.NewCustomerClient(custConn)

	r := mux.NewRouter()
	routes.RegisterCustomerRoutes(r, c)
	log.Fatal(http.ListenAndServe(apiGatewayPort, r))

}
