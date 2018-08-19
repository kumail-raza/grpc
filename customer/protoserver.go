package customer

import (
	"context"

	"github.com/globalsign/mgo/bson"
	"github.com/minhajuddinkhan/grpc/db"
	"github.com/minhajuddinkhan/grpc/protocols"
)

//Server Server
type Server struct {
	Database *db.Mongo
}

//GetCustomer GetCustomer
func (s *Server) GetCustomer(ctx context.Context, customerProto *customerprotocol.Id) (*customerprotocol.CustomerProto, error) {

	collection := s.Database.GetCollection("customers")
	var customer Customer
	collection.FindId(bson.ObjectIdHex(customerProto.GetId())).One(&customer)
	return &customerprotocol.CustomerProto{
		Id: func(str string) *string {
			return &str
		}(customer.ID.Hex()),
		Name:    &customer.Name,
		Address: &customer.Address,
		Email:   &customer.Email,
	}, nil
}

//CreateCustomer CreateCustomer
func (s *Server) CreateCustomer(ctx context.Context, customerProto *customerprotocol.CustomerProto) (*customerprotocol.Id, error) {

	collection := s.Database.GetCollection("customers")
	cust := NewCustomer(*customerProto.Name, *customerProto.Address, *customerProto.Address)
	err := collection.Insert(cust)
	if err != nil {
		return nil, err
	}

	customerID := cust.ID.Hex()
	return &customerprotocol.Id{Id: &customerID}, nil
}

//GetAllCustomers GetAllCustomers
func (s *Server) GetAllCustomers(ctx context.Context, in *customerprotocol.NothingFancy) (*customerprotocol.CustomersProto, error) {

	collection := s.Database.GetCollection("customers")
	var customers []Customer
	collection.Find(bson.M{}).All(&customers)

	customersProto := customerprotocol.CustomersProto{}
	for _, c := range customers {
		customersProto.Customers = append(customersProto.Customers, &customerprotocol.CustomerProto{
			Name:    &c.Name,
			Address: &c.Address,
			Email:   &c.Email,
			Id: func(str string) *string {
				return &str
			}(c.ID.Hex()),
		})
	}

	return &customersProto, nil
}
