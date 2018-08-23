package customer

import (
	"context"

	"github.com/minhajuddinkhan/grpc/convert"

	"github.com/globalsign/mgo/bson"
	"github.com/minhajuddinkhan/grpc/db"
)

//Server Server
type Server struct {
	Database *db.Mongo
}

//GetCustomer GetCustomer
func (s *Server) GetCustomer(ctx context.Context, customerProto *Id) (*CustomerProto, error) {

	collection := s.Database.GetCollection("customers")
	var customer Customer
	collection.FindId(bson.ObjectIdHex(customerProto.GetId())).One(&customer)

	cPro := ToProtoCustomer(customer)
	return &cPro, nil
}

//CreateCustomer CreateCustomer
func (s *Server) CreateCustomer(ctx context.Context, customerProto *CustomerProto) (*Id, error) {

	collection := s.Database.GetCollection("customers")
	cust := NewCustomer(*customerProto.Name, *customerProto.Address, *customerProto.Address)
	err := collection.Insert(cust)
	if err != nil {
		return nil, err
	}
	return &Id{Id: convert.MongoIDToStringPtr(cust.ID)}, nil
}

//GetAllCustomers GetAllCustomers
func (s *Server) GetAllCustomers(ctx context.Context, in *NothingFancy) (*CustomersProto, error) {

	collection := s.Database.GetCollection("customers")
	var customers []Customer
	collection.Find(bson.M{}).All(&customers)

	customersProto := ToMultipleProtoCustomer(customers)
	return &customersProto, nil
}
