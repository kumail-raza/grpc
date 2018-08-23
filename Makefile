build:
	protoc -I. --go_out=plugins=grpc:$(GOPATH)/src/github.com/minhajuddinkhan/grpc \
	  customer/customer.proto
	protoc -I. --go_out=plugins=grpc:$(GOPATH)/src/github.com/minhajuddinkhan/grpc \
	  order/order.proto	