package grpcproxyserver

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/xaionaro-go/grpcproxy/protobuf/go/proxy_grpc"
)

type GRPCProxyServer struct {
	NextConnectionID atomic.Uint64
	ConnectionMap    sync.Map

	proxy_grpc.UnimplementedNetworkProxyServer
}

var _ proxy_grpc.NetworkProxyServer = (*GRPCProxyServer)(nil)

func New() *GRPCProxyServer {
	return &GRPCProxyServer{}
}

func (srv *GRPCProxyServer) Proxy(proxyServer proxy_grpc.NetworkProxy_ProxyServer) error {
	msg, err := proxyServer.Recv()
	if err != nil {
		return fmt.Errorf("unable to get the connection request message: %w", err)
	}

	connectMsg, ok := msg.GetMessageForwardOneOf().(*proxy_grpc.MessageForward_Connect)
	if !ok {
		return fmt.Errorf("the first message is expected to be %T, but received %T", connectMsg, msg.GetMessageForwardOneOf())
	}

	connectReq := connectMsg.Connect

	msg.GetMessageForwardOneOf()
	conn := newConnection(
		srv.nextConnectionID(),
		connectReq.GetProtocol(),
		connectReq.GetHostname(),
		uint16(connectReq.GetPort()),
	)
	return conn.serve(proxyServer.Context(), proxyServer)
}

func (srv *GRPCProxyServer) nextConnectionID() connectionID {
	return connectionID(srv.NextConnectionID.Add(1))
}
