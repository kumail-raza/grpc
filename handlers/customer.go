package handlers

import (
	"context"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"

	boom "github.com/darahayes/go-boom"
	"github.com/globalsign/mgo/bson"
	"github.com/minhajuddinkhan/grpc/customer"
	"github.com/minhajuddinkhan/todogo/utils"
)

//GetCustomers GetCustomers
func GetCustomers(c customer.CustomerClient) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		cID := mux.Vars(r)["id"]
		logrus.Info(cID)
		incomingCustomer, err := c.GetCustomer(context.Background(), &customer.Id{Id: &cID})
		if err != nil {
			boom.BadRequest(w, err)
			return
		}

		logrus.Info(*incomingCustomer.Id)
		utils.Respond(w, customer.Customer{
			ID:      bson.ObjectIdHex(*incomingCustomer.Id),
			Name:    *incomingCustomer.Name,
			Address: *incomingCustomer.Address,
			Email:   *incomingCustomer.Email,
		})
	}

}

//CreateCustomer CreateCustomer
func CreateCustomer(c customer.CustomerClient) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var cust customer.Customer
		err := utils.DecodeRequest(r, &cust)
		if err != nil {
			boom.BadRequest(w, err)
			return
		}
		custID := string(bson.NewObjectId())
		customerProtocol := &customer.CustomerProto{
			Id:      &custID,
			Name:    &cust.Name,
			Address: &cust.Address,
			Email:   &cust.Email,
		}
		newCustomerIDProto, err := c.CreateCustomer(r.Context(), customerProtocol)
		if err != nil {
			boom.BadRequest(w, err)
			return
		}
		utils.Respond(w, struct {
			CustomerID string
		}{
			*newCustomerIDProto.Id,
		})
		return

	}
}

//GetAllCustomers GetAllCustomers
func GetAllCustomers(c customer.CustomerClient) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		customersProto, err := c.GetAllCustomers(r.Context(), &customer.NothingFancy{})
		if err != nil {
			boom.BadRequest(w, err)
			return
		}
		var customers []customer.Customer
		for _, cust := range customersProto.Customers {
			customers = append(customers, customer.Customer{
				ID:      bson.ObjectIdHex(*cust.Id),
				Name:    *cust.Name,
				Address: *cust.Address,
				Email:   *cust.Email,
			})
		}

		utils.Respond(w, customers)

	}

}
