package order

import (
	"context"
	"errors"

	"github.com/minhajuddinkhan/grpc/convert"

	"github.com/sirupsen/logrus"

	"github.com/globalsign/mgo/bson"
	"github.com/minhajuddinkhan/grpc/db"
)

//Server Server
type Server struct {
	Database *db.Mongo
}

//GetOrder GetOrder
func (s *Server) GetOrder(ctx context.Context, in *Id) (*OrderProto, error) {
	collection := s.Database.GetCollection("orders")
	var order Order

	query := []bson.M{{"$lookup": bson.M{ // lookup the documents table here
		"from":         "customers",
		"localField":   "customerid",
		"foreignField": "_id",
		"as":           "customer",
	}},
		{"$match": bson.M{"_id": convert.StringPtrToMongoID(in.Id)}},
	}
	if err := collection.Pipe(query).One(&order); err != nil {
		return nil, err
	}

	customerProtos := ToMultipleProtoCustomer(order.Customer)
	ordersProto := &OrderProto{
		Id:         convert.MongoIDToStringPtr(order.ID),
		Name:       &order.Name,
		Address:    &order.Address,
		CustomerId: convert.MongoIDToStringPtr(order.CustomerID),
		Customers:  &customerProtos,
	}

	return ordersProto, nil
}

//CreateOrder CreateOrder
func (s *Server) CreateOrder(ctx context.Context, in *OrderProto) (*Id, error) {

	collection := s.Database.GetCollection("orders")
	logrus.Info(in)
	order := Order{
		Address:    *in.Address,
		CustomerID: bson.ObjectIdHex(*in.CustomerId),
		Name:       *in.Name,
	}
	err := collection.Insert(order)
	if err != nil {
		return nil, err
	}
	orderID := order.ID.Hex()
	return &Id{Id: &orderID}, nil

}

//GetAllOrders GetAllOrders
func (s *Server) GetAllOrders(ctx context.Context, in *Noop) (*OrdersProto, error) {

	var orders []Order

	query := []bson.M{{"$lookup": bson.M{ // lookup the documents table here
		"from":         "customers",
		"localField":   "customerid",
		"foreignField": "_id",
		"as":           "customer",
	}}}

	if err := s.Database.GetCollection("orders").Pipe(query).All(&orders); err != nil {
		return nil, err
	}
	if len(orders) == 0 {
		return nil, errors.New("No Orders Yet")
	}

	result := OrdersProto{}
	for _, o := range orders {
		customerProtos := ToMultipleProtoCustomer(o.Customer)
		result.Orders = append(result.Orders, &OrderProto{
			Name:       &o.Name,
			Address:    &o.Address,
			CustomerId: convert.MongoIDToStringPtr(o.CustomerID),
			Id:         convert.MongoIDToStringPtr(o.ID),
			Customers:  &customerProtos,
		})
	}
	return &result, nil

}
