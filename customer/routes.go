package customer

import (
	"github.com/gorilla/mux"
	"github.com/minhajuddinkhan/grpc-test/db"
)

//NewCustomerRouter NewCustomerRouter
func NewCustomerRouter(mongo *db.Mongo) *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/customers", GetAllCustomers(mongo)).Methods("GET")
	r.HandleFunc("/customers", CreateCustomer(mongo)).Methods("POST")
	return r
}
