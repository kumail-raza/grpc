package routes

import (
	"github.com/gorilla/mux"
	"github.com/minhajuddinkhan/grpc/handlers"
	"github.com/minhajuddinkhan/grpc/protocols"
)

//RegisterCustomerRoutes RegisterCustomerRoutes
func RegisterCustomerRoutes(r *mux.Router, c customerprotocol.CustomerClient) *mux.Router {

	r.HandleFunc("/customers/{id}", handlers.GetCustomers(c)).Methods("GET")
	r.HandleFunc("/customers", handlers.GetAllCustomers(c)).Methods("GET")
	r.HandleFunc("/customers", handlers.CreateCustomer(c)).Methods("POST")
	return r
}
