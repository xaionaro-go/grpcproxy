package grpchttpproxy

import (
	"context"
	"fmt"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/facebookincubator/go-belt/tool/logger"
	"github.com/xaionaro-go/grpcproxy/protobuf/go/proxy_grpc"
)

type ProxiedConn struct {
	ctx              context.Context
	connectionID     atomic.Uint64
	connectionStatus atomic.Uint32
	cancelFn         context.CancelFunc
	waitGroup        sync.WaitGroup
	proxyClient      proxy_grpc.NetworkProxyClient
	errorOnce        sync.Once
	error            error
	connToProxy      proxy_grpc.NetworkProxy_ProxyClient
	protocol         proxy_grpc.NetworkProtocol
	remoteAddr       net.Addr
	metadataLocker   sync.Mutex

	receiveLocker sync.Mutex
	receiveBuffer []byte
}

var _ net.Conn = (*ProxiedConn)(nil)

func (d *Dialer) newConn(
	ctx context.Context,
	proto proxy_grpc.NetworkProtocol,
	host string,
	port uint16,
) (net.Conn, error) {
	return newConn(ctx, d.NetworkProxyClient, proto, host, port)
}

func newConn(
	ctx context.Context,
	proxyClient proxy_grpc.NetworkProxyClient,
	proto proxy_grpc.NetworkProtocol,
	host string,
	port uint16,
) (net.Conn, error) {
	ctx, cancelFn := context.WithCancel(ctx)
	conn := &ProxiedConn{
		ctx:         ctx,
		proxyClient: proxyClient,
		protocol:    proto,
		remoteAddr:  newAddr(proto, net.ParseIP(host), port),
		cancelFn:    cancelFn,
	}
	err := conn.connect(ctx, proto, host, port)
	if err != nil {
		conn.close()
		conn.waitGroup.Wait()
		return nil, fmt.Errorf("unable to connect to %s://%s:%d: %w", proto, host, port, err)
	}
	return conn, nil
}

func (conn *ProxiedConn) closeWithError(err error) {
	conn.errorOnce.Do(func() {
		conn.error = err
	})
	conn.close()
}

func (conn *ProxiedConn) setConnectionStatus(
	ctx context.Context,
	connStatus proxy_grpc.ConnectionStatus,
) {
	logger.Debugf(ctx, "new connection status: %v", connStatus)
	conn.metadataLocker.Lock()
	defer conn.metadataLocker.Unlock()

	conn.connectionStatus.Store(uint32(connStatus))
	switch connStatus {
	case proxy_grpc.ConnectionStatus_Connecting:
	case proxy_grpc.ConnectionStatus_Connected:
	case proxy_grpc.ConnectionStatus_Rejected:
		conn.closeWithError(fmt.Errorf("the connection was rejected"))
	case proxy_grpc.ConnectionStatus_Finishing:
	case proxy_grpc.ConnectionStatus_Finished:
		conn.close()
	case proxy_grpc.ConnectionStatus_Reset:
		conn.closeWithError(fmt.Errorf("the connection was reset"))
	case proxy_grpc.ConnectionStatus_Unreachable:
		conn.closeWithError(fmt.Errorf("destination unreachable"))
	case proxy_grpc.ConnectionStatus_TimedOut:
		conn.closeWithError(fmt.Errorf("connection timed out"))
	case proxy_grpc.ConnectionStatus_UnableToResolve:
		conn.closeWithError(fmt.Errorf("unable to resolve"))
	default:
		conn.closeWithError(fmt.Errorf("unexpected connection status: %v", connStatus))
	}
}

func (conn *ProxiedConn) setRemoteAddr(
	_ context.Context,
	proto proxy_grpc.NetworkProtocol,
	ipAddr []byte,
	port uint16,
) {
	conn.metadataLocker.Lock()
	defer conn.metadataLocker.Unlock()
	conn.remoteAddr = newAddr(proto, ipAddr, port)
}

func (conn *ProxiedConn) setConnectionID(
	ctx context.Context,
	newConnID uint64,
) {
	logger.Debugf(ctx, "new connection ID: %d", newConnID)
	conn.metadataLocker.Lock()
	defer conn.metadataLocker.Unlock()

	// actually connection ID is set only once, and all other times
	// it is rewritten with the exact same value.
	conn.connectionID.Store(newConnID)
}

func (conn *ProxiedConn) waitForConnect(
	ctx context.Context,
	connToProxy proxy_grpc.NetworkProxy_ProxyClient,
) error {
	logger.Debugf(ctx, "waitForConnect")
	defer logger.Debugf(ctx, "waitForConnect")

	for {
		msg, err := connToProxy.Recv()
		if err != nil {
			return fmt.Errorf("unable to receive a message from the proxy: %w", err)
		}

		var connUpdate *proxy_grpc.ConnectionStatusUpdate
		switch oneof := msg.MessageBackOneOf.(type) {
		case *proxy_grpc.MessageBack_ConnectionUpdate:
			connUpdate = oneof.ConnectionUpdate
		case *proxy_grpc.MessageBack_ReceivedPayload:
			return fmt.Errorf("received a traffic payload before the connection was established")
		default:
			return fmt.Errorf("receive a message of type %T, while expecting %T", oneof, (*proxy_grpc.MessageBack_ConnectionUpdate)(nil))
		}

		switch s := connUpdate.GetConnectionStatus(); s {
		case proxy_grpc.ConnectionStatus_Connecting:
			continue
		case proxy_grpc.ConnectionStatus_Connected:
			conn.setConnectionID(ctx, connUpdate.ConnectionID)
			conn.setConnectionStatus(ctx, connUpdate.ConnectionStatus)
			if connUpdate.RemotePort != nil {
				conn.setRemoteAddr(ctx, conn.protocol, connUpdate.RemoteIPAddress, uint16(*connUpdate.RemotePort))
			}
			return nil
		case proxy_grpc.ConnectionStatus_Rejected:
			return fmt.Errorf("rejected")
		case proxy_grpc.ConnectionStatus_Finishing:
			return fmt.Errorf("unexpected: finishing")
		case proxy_grpc.ConnectionStatus_Finished:
			return fmt.Errorf("unexpected: finished")
		case proxy_grpc.ConnectionStatus_Reset:
			return fmt.Errorf("reset")
		case proxy_grpc.ConnectionStatus_Unreachable:
			return fmt.Errorf("unreachable")
		case proxy_grpc.ConnectionStatus_TimedOut:
			return fmt.Errorf("timed out")
		case proxy_grpc.ConnectionStatus_UnableToResolve:
			return fmt.Errorf("unable to resolve")
		case proxy_grpc.ConnectionStatus_BrokenProxyConnection:
			return fmt.Errorf("broken connection to the proxy")
		default:
			return fmt.Errorf("unexpected connection status: %v", s)
		}
	}

}

func (conn *ProxiedConn) connect(
	ctx context.Context,
	proto proxy_grpc.NetworkProtocol,
	host string,
	port uint16,
) (_err error) {
	logger.Tracef(ctx, "connect")
	defer func() { logger.Tracef(ctx, "/connect: %v", _err) }()

	connToProxy, err := conn.proxyClient.Proxy(ctx)
	if err != nil {
		return fmt.Errorf("unable to establish the connection with the proxy server: %w", err)
	}

	err = connToProxy.Send(&proxy_grpc.MessageForward{
		MessageForwardOneOf: &proxy_grpc.MessageForward_Connect{
			Connect: &proxy_grpc.ConnectRequest{
				Protocol: proto,
				Hostname: host,
				Port:     uint32(port),
			},
		},
	})
	if err != nil {
		return fmt.Errorf("unable to request the proxy connecting to %s://%s:%d: %w", proto, host, port, err)
	}

	err = conn.waitForConnect(ctx, connToProxy)
	if err != nil {
		return fmt.Errorf("connection to %s://%s:%d failed: %w", proto, host, port, err)
	}

	conn.connToProxy = connToProxy
	return nil
}

func (conn *ProxiedConn) Read(b []byte) (int, error) {
	conn.waitGroup.Add(1)
	defer conn.waitGroup.Done()

	conn.receiveLocker.Lock()
	defer conn.receiveLocker.Unlock()

	if conn.isClosed() {
		return 0, fmt.Errorf("connection is closed")
	}

	for {
		copy(b, conn.receiveBuffer)
		nRead := min(len(b), len(conn.receiveBuffer))
		conn.receiveBuffer = conn.receiveBuffer[nRead:]
		if nRead > 0 {
			logger.Tracef(conn.ctx, "received %d bytes", nRead)
			return nRead, nil
		}
		logger.Tracef(conn.ctx, "waiting for incoming data (requested buffer size: %d)...", len(b))
		b = b[nRead:]

		msg, err := conn.connToProxy.Recv()
		if err != nil {
			return 0, fmt.Errorf("unable to read a proxied packet: %w", err)
		}

		var payload []byte
		switch m := msg.GetMessageBackOneOf().(type) {
		case *proxy_grpc.MessageBack_ConnectionUpdate:
			return 0, fmt.Errorf("received a message of type %T, while expected %T", msg.GetMessageBackOneOf(), m)
		case *proxy_grpc.MessageBack_ReceivedPayload:
			payload = m.ReceivedPayload
		default:
			return 0, fmt.Errorf("received a message of an unexpected type %T", m)
		}

		conn.receiveBuffer = append(conn.receiveBuffer, payload...)
	}
}

func (conn *ProxiedConn) Write(b []byte) (int, error) {
	conn.waitGroup.Add(1)
	defer conn.waitGroup.Done()

	if conn.isClosed() {
		return 0, fmt.Errorf("connection is closed")
	}
	logger.Tracef(conn.ctx, "sending %d bytes", len(b))
	err := conn.connToProxy.Send(&proxy_grpc.MessageForward{
		MessageForwardOneOf: &proxy_grpc.MessageForward_SendPayload{
			SendPayload: b,
		},
	})
	if err != nil {
		return 0, fmt.Errorf("unable to send the data via the proxy: %w", err)
	}
	return len(b), nil
}

func (conn *ProxiedConn) isClosed() bool {
	select {
	case <-conn.ctx.Done():
		return true
	default:
		return false
	}
}

func (conn *ProxiedConn) close() {
	conn.cancelFn()
}

func (conn *ProxiedConn) Close() error {
	conn.close()
	conn.waitGroup.Wait()
	return nil
}

func (conn *ProxiedConn) LocalAddr() net.Addr {
	conn.metadataLocker.Lock()
	defer conn.metadataLocker.Unlock()
	return nil
}

func (conn *ProxiedConn) RemoteAddr() net.Addr {
	conn.metadataLocker.Lock()
	defer conn.metadataLocker.Unlock()
	return conn.remoteAddr
}

func (conn *ProxiedConn) SetDeadline(t time.Time) error {
	return fmt.Errorf("not implemented, yet")
}

func (conn *ProxiedConn) SetReadDeadline(t time.Time) error {
	return fmt.Errorf("not implemented, yet")
}

func (conn *ProxiedConn) SetWriteDeadline(t time.Time) error {
	return fmt.Errorf("not implemented, yet")
}
