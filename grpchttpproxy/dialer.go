package grpchttpproxy

import (
	"context"
	"fmt"
	"net"
	"strconv"

	"github.com/xaionaro-go/grpcproxy/protobuf/go/proxy_grpc"
)

// Dialer is similar to net.Dialer, but it creates connection proxied
// via the gRPC server (see package `grpcproxyserver` for the server side).
type Dialer struct {
	proxy_grpc.NetworkProxyClient
}

// NewDialer returns a new instance of Dialer.
func NewDialer(
	proxy proxy_grpc.NetworkProxyClient,
) *Dialer {
	return &Dialer{
		NetworkProxyClient: proxy,
	}
}

// DialContext is similar to (*net.Dialer).DialContext, but the created
// connection is proxied via the gRPC server (see package `grpcproxyserver` for the server side).
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
