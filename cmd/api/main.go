package main

import (
	"log"
	"net/http"

	"github.com/minhajuddinkhan/grpc/order"

	"github.com/gorilla/mux"
	"github.com/minhajuddinkhan/grpc/customer"
	"github.com/minhajuddinkhan/grpc/routes"
	"google.golang.org/grpc"
)

const (
	apiGatewayPort           = ":3000"
	customerMicroServiceAddr = "localhost:3443"
	orderMicroServiceAddr    = "localhost:4443"
)

func main() {

	custConn, err := grpc.Dial(customerMicroServiceAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer custConn.Close()

	orderConn, err := grpc.Dial(orderMicroServiceAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect to order service : %v", err)
	}
	defer orderConn.Close()

	c := customer.NewCustomerClient(custConn)
	o := order.NewOrderClient(orderConn)

	r := mux.NewRouter()
	routes.RegisterCustomerRoutes(r, c)
	routes.RegisterOrderRoutes(r, o)
	log.Fatal(http.ListenAndServe(apiGatewayPort, r))

}
