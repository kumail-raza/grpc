build:
	protoc -I. --go_out=plugins=grpc:$(GOPATH)/src/github.com/minhajuddinkhan/grpc \
	  protocols/customer.proto