package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/minhajuddinkhan/todogo/utils"

	"github.com/darahayes/go-boom"
	"github.com/globalsign/mgo/bson"

	"github.com/minhajuddinkhan/grpc-test/customer"
	"github.com/minhajuddinkhan/grpc-test/customer/proto"

	"google.golang.org/grpc"
)

//GetCustomers GetCustomers
func GetCustomers(w http.ResponseWriter, r *http.Request) {

	address := "localhost:3443"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := customerproto.NewCustomerClient(conn)

	cID := "1"
	incomingCustomer, err := c.GetCustomer(context.Background(), &customerproto.Id{Id: &cID})
	if err != nil {
		boom.BadRequest(w, err)
		return
	}
	utils.Respond(w, customer.Customer{
		ID:      bson.ObjectIdHex(*incomingCustomer.Id),
		Name:    *incomingCustomer.Name,
		Address: *incomingCustomer.Address,
		Email:   *incomingCustomer.Email,
	})

}
