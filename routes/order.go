package routes

import (
	"github.com/gorilla/mux"
	"github.com/minhajuddinkhan/grpc/handlers"
	"github.com/minhajuddinkhan/grpc/order"
)

//RegisterOrderRoutes RegisterOrderRoutes
func RegisterOrderRoutes(r *mux.Router, o order.OrderClient) *mux.Router {

	r.HandleFunc("/orders", handlers.GetAllOrders(o)).Methods("GET")
	r.HandleFunc("/orders/{id}", handlers.GetOrder(o)).Methods("GET")
	r.HandleFunc("/orders", handlers.CreateOrder(o)).Methods("POST")
	return r
}
