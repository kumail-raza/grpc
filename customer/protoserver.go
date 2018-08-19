package customer

import (
	"context"

	"github.com/globalsign/mgo/bson"
	"github.com/minhajuddinkhan/grpc-test/customer/proto"
	"github.com/minhajuddinkhan/grpc-test/db"
)

//Server Server
type Server struct {
	Database *db.Mongo
}

//GetCustomer GetCustomer
func (s *Server) GetCustomer(ctx context.Context, customerProto *customerproto.Id) (*customerproto.CustomerProto, error) {

	collection := s.Database.GetCollection("customer")
	var customer Customer
	collection.Find(bson.M{"_id": customerProto.GetId()}).One(&customer)
	cID := string(customer.ID)
	return &customerproto.CustomerProto{
		Id:      &cID,
		Name:    &customer.Name,
		Address: &customer.Address,
		Email:   &customer.Email,
	}, nil
}
