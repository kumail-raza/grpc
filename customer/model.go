package customer

import "github.com/globalsign/mgo/bson"

//Customer Customer
type Customer struct {
	ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name    string
	Address string
	Email   string
}

//NewCustomer NewCustomer
func NewCustomer(name, address, email string) Customer {
	return Customer{
		ID:      bson.NewObjectId(),
		Name:    name,
		Address: address,
		Email:   email,
	}
}
