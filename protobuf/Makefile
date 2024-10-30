
all: go

go:
	protoc --go_out=. --go-grpc_out=. proxy.proto
	mv github.com/xaionaro-go/grpcproxy/protobuf/go .
	rm -rf github.com

