package grpchttpproxy

import (
	"github.com/xaionaro-go/grpcproxy/protobuf/go/proxy_grpc"
)

func NetworkToProtocol(network string) proxy_grpc.NetworkProtocol {
	switch network {
	case "tcp", "tcp4", "tcp6":
		return proxy_grpc.NetworkProtocol_TCP
	default:
		return proxy_grpc.NetworkProtocol_networkProtocolUndefined
	}
}
