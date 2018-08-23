package handlers

import (
	"context"
	"net/http"

	"github.com/minhajuddinkhan/grpc/customer"

	"github.com/minhajuddinkhan/grpc/convert"

	"github.com/darahayes/go-boom"
	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
	"github.com/minhajuddinkhan/grpc/order"
	"github.com/minhajuddinkhan/todogo/utils"
)

//GetAllOrders  GetAllOrders
func GetAllOrders(o order.OrderClient) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		ordersProto, err := o.GetAllOrders(r.Context(), &order.Noop{})
		if err != nil {
			boom.BadRequest(w, err)
			return
		}
		var os []order.Order
		for _, orderIteree := range ordersProto.Orders {
			os = append(os, order.Order{
				ID:         bson.ObjectIdHex(*orderIteree.Id),
				Name:       *orderIteree.Name,
				Address:    *orderIteree.Address,
				CustomerID: convert.StringPtrToMongoID(orderIteree.CustomerId),
				Customer: func(cpros *order.CustomersProto) []customer.Customer {
					customers := []customer.Customer{}
					for _, c := range cpros.Customers {
						customers = append(customers, customer.Customer{
							ID:      convert.StringPtrToMongoID(c.Id),
							Name:    *c.Name,
							Address: *c.Address,
							Email:   *c.Email,
						})
					}
					return customers
				}(orderIteree.Customers),
			})
		}
		utils.Respond(w, os)
	}
}

//GetOrder GetOrder
func GetOrder(o order.OrderClient) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		oID := mux.Vars(r)["id"]
		o, err := o.GetOrder(context.Background(), &order.Id{Id: &oID})
		if err != nil {
			boom.BadRequest(w, err)
			return
		}

		utils.Respond(w, order.Order{
			ID:         bson.ObjectIdHex(*o.Id),
			Name:       *o.Name,
			Address:    *o.Address,
			CustomerID: convert.StringPtrToMongoID(o.CustomerId),
			Customer: func(cpros *order.CustomersProto) []customer.Customer {
				customers := []customer.Customer{}
				for _, c := range cpros.Customers {
					customers = append(customers, customer.Customer{
						ID:      convert.StringPtrToMongoID(c.Id),
						Name:    *c.Name,
						Address: *c.Address,
						Email:   *c.Email,
					})
				}
				return customers
			}(o.Customers),
		})

	}
}

//CreateOrder CreateOrder
func CreateOrder(o order.OrderClient) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var ord order.Order
		err := utils.DecodeRequest(r, &ord)
		if err != nil {
			boom.BadRequest(w, err)
			return
		}

		orderID := string(bson.NewObjectId())
		orderProtocol := &order.OrderProto{
			Id:         &orderID,
			Name:       &ord.Name,
			Address:    &ord.Address,
			CustomerId: convert.MongoIDToStringPtr(ord.CustomerID),
		}
		newOrderIDProtocol, err := o.CreateOrder(r.Context(), orderProtocol)
		if err != nil {
			boom.BadRequest(w, err)
			return
		}
		utils.Respond(w, struct {
			CustomerID string
		}{
			*newOrderIDProtocol.Id,
		})
	}
}
