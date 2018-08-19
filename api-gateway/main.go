package main

import (
	"github.com/gorilla/mux"
	"github.com/minhajuddinkhan/grpc-test/handlers"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/customer", handlers.GetCustomers)
}
