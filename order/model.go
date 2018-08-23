package order

import (
	"github.com/globalsign/mgo/bson"
	"github.com/minhajuddinkhan/grpc/customer"
)

//Order Order
type Order struct {
	ID         bson.ObjectId       `json:"id" bson:"_id,omitempty"`
	Name       string              `json:"name"`
	Address    string              `json:"address"`
	CustomerID bson.ObjectId       `json:"customerId"`
	Customer   []customer.Customer `json:"customer"`
}
