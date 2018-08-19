package order

import "github.com/globalsign/mgo/bson"

//Order Order
type Order struct {
	ID         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	CustomerID bson.ObjectId `json:"customerId" bson:"_id,omitempty"`
	Name       string
}
