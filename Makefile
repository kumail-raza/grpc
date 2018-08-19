build:
	protoc -I. --go_out=plugins=grpc:$(GOPATH)/src/github.com/minhajuddinkhan/grpc-test \
	  api/proto/api.proto