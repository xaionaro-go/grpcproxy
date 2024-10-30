# About

[![Go Reference](https://godoc.org/github.com/xaionaro-go/grpcproxy?status.svg)](https://godoc.org/github.com/xaionaro-go/grpcproxy)
[![Go Report Card](https://goreportcard.com/badge/github.com/xaionaro-go/grpcproxy?branch=main)](https://goreportcard.com/report/github.com/xaionaro-go/grpcproxy)

This is a gRPC server and a client that allows to proxy network connections via the server.

# How to use

## Server-side

Let's say you already have a gRPC server written in Go, something like this:
```go
	grpcServer := grpc.NewServer()
	myFancyServiceGRPC := server.NewGRPCServer(myFancyService)
	myFancyService_grpc.RegisterMyFancyServiceServer(grpcServer, myFancyServiceGRPC)
```

To add the proxy capability you just need to register the proxy service as well:
```go
import "github.com/xaionaro-go/grpcproxy/grpcproxyserver"
import "github.com/xaionaro-go/grpcproxy/protobuf/go/proxy_grpc"

	grpcServer := grpc.NewServer()
	myFancyService := myfancyservice.NewServer(myFancyService)
	myfancyservice_grpc.RegisterMyFancyServiceServer(grpcServer, myFancyService)

	proxyServer := grpcproxyserver.New()
	proxy_grpc.RegisterNetworkProxyServer(grpcServer, proxyServer)
```

## Client-side

On the client side you may replace a direct connection:
```go
conn, err := net.Dial("tcp", addr)
```
with proxied connection:
```go
conn, err := grpchttpproxy.NewDialer(myGRPCClient).DialContext(ctx, "tcp", "addr")
```

Or if you need to proxy an HTTP request, you may make an HTTP client:
```go
	httpClient := &http.Client{
		Transport: &http.Transport{
			DialContext: grpchttpproxy.NewDialer(myGRPCClient).DialContext,
		},
	}
```
And then perform the HTTP requests via this client
