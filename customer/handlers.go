package customer

import (
	"net/http"

	"github.com/globalsign/mgo/bson"
	"github.com/minhajuddinkhan/todogo/utils"

	"github.com/darahayes/go-boom"

	"github.com/minhajuddinkhan/grpc-test/db"
)

//CreateCustomer CreateCustomer
func CreateCustomer(mongo *db.Mongo) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var customer Customer
		if err := utils.DecodeRequest(r, &customer); err != nil {
			boom.BadRequest(w, err)
			return
		}
		customer.ID = bson.NewObjectId()
		if err := mongo.GetCollection("customers").Insert(&customer); err != nil {
			boom.BadData(w, err)
		}
		utils.Respond(w, customer)
		return
	}

}

//GetAllCustomers GetAllCustomers
func GetAllCustomers(mongo *db.Mongo) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var customers []Customer
		err := mongo.GetCollection("customers").Find(bson.M{}).All(&customers)
		if err != nil {
			boom.BadRequest(w, err)
			return
		}
		utils.Respond(w, customers)
		return

	}

}
