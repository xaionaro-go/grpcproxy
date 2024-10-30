package grpchttpproxy

import (
	"net"

	"github.com/xaionaro-go/grpcproxy/protobuf/go/proxy_grpc"
)

func newAddr(
	proto proxy_grpc.NetworkProtocol,
	ipAddr net.IP,
	port uint16,
) net.Addr {
	switch proto {
	case proxy_grpc.NetworkProtocol_TCP:
		return &net.TCPAddr{
			IP:   ipAddr,
			Port: int(port),
		}
	default:
		return nil
	}
}
