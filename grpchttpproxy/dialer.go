package grpchttpproxy

import (
	"context"
	"fmt"
	"net"
	"strconv"

	"github.com/xaionaro-go/grpcproxy/protobuf/go/proxy_grpc"
)

type Dialer struct {
	proxy_grpc.NetworkProxyClient
}

func NewDialer(
	proxy proxy_grpc.NetworkProxyClient,
) *Dialer {
	return &Dialer{
		NetworkProxyClient: proxy,
	}
}

func (d *Dialer) DialContext(
	ctx context.Context,
	network string,
	addr string,
) (net.Conn, error) {
	host, portString, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, fmt.Errorf("unable to split '%s' to host and port: %w", addr, err)
	}

	port, err := strconv.ParseUint(portString, 10, 16)
	if err != nil {
		return nil, fmt.Errorf("unable to parse port '%s': %w", portString, err)
	}

	proto := NetworkToProtocol(network)
	if proto == proxy_grpc.NetworkProtocol_networkProtocolUndefined {
		return nil, fmt.Errorf("unknown network protocol '%s'", network)
	}

	return d.newConn(ctx, proto, host, uint16(port))
}
