build:
	protoc -I. --go_out=plugins=grpc:$(GOPATH)/src/github.com/minhajuddinkhan/grpc-test \
	  customer/proto/customer.proto