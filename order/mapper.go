package order

import (
	"github.com/minhajuddinkhan/grpc/convert"
	"github.com/minhajuddinkhan/grpc/customer"
)

//ProtoCustomerToCustomerStruct ProtoCustomerToCustomerStruct
func ProtoCustomerToCustomerStruct(c CustomerProto) customer.Customer {
	customer := customer.Customer{
		Name:    *c.Name,
		Address: *c.Address,
		Email:   *c.Email,
		ID:      convert.StringPtrToMongoID(c.Id),
	}
	return customer
}

//ToProtoCustomer ToProtoCustomer
func ToProtoCustomer(c customer.Customer) CustomerProto {
	customerProto := CustomerProto{
		Id:      convert.MongoIDToStringPtr(c.ID),
		Name:    &c.Name,
		Address: &c.Address,
		Email:   &c.Email,
	}
	return customerProto
}

//MultipleProtoCustomerToCustomerStruct MultipleProtoCustomerToCustomerStruct
func MultipleProtoCustomerToCustomerStruct(c CustomersProto) []customer.Customer {

	customers := []customer.Customer{}
	for _, m := range c.Customers {
		x := ProtoCustomerToCustomerStruct(*m)
		customers = append(customers, x)
	}
	return customers
}

//ToMultipleProtoCustomer ToMultipleProtoCustomer
func ToMultipleProtoCustomer(c []customer.Customer) CustomersProto {

	customersProto := CustomersProto{}
	for _, m := range c {
		x := ToProtoCustomer(m)
		customersProto.Customers = append(customersProto.Customers, &x)
	}
	return customersProto
}
