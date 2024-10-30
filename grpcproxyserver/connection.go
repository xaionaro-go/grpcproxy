package grpcproxyserver

import (
	"context"
	"fmt"
	"net"
	"sync"

	"github.com/facebookincubator/go-belt/tool/logger"
	"github.com/hashicorp/go-multierror"
	"github.com/xaionaro-go/grpcproxy/protobuf/go/proxy_grpc"
)

type connectionID uint64

type connectionStatus struct {
	Err    error
	Status proxy_grpc.ConnectionStatus
}

type connection struct {
	locker                 sync.Mutex
	waitGroup              sync.WaitGroup
	id                     connectionID
	proto                  proxy_grpc.NetworkProtocol
	host                   string
	port                   uint16
	cancelFunc             context.CancelFunc
	statusChan             chan connectionStatus
	remoteAddr             net.Addr
	conn                   net.Conn
	readerTargetBuffer     [65536]byte
	readTargetBuffer       [65536]byte
	readTargetBufferLocker sync.Mutex
	receivedFromTarget     chan []byte
	receivedFromClient     chan []byte
}

func newConnection(
	connectionID connectionID,
	proto proxy_grpc.NetworkProtocol,
	host string,
	port uint16,
) *connection {
	return &connection{
		id:    connectionID,
		proto: proto,
		host:  host,
		port:  port,

		statusChan:         make(chan connectionStatus),
		receivedFromTarget: make(chan []byte),
		receivedFromClient: make(chan []byte),
	}
}

func (conn *connection) GetID() connectionID {
	conn.locker.Lock()
	defer conn.locker.Unlock()
	return conn.id
}

func (conn *connection) GetRemoteAddr() net.Addr {
	conn.locker.Lock()
	defer conn.locker.Unlock()
	return conn.remoteAddr
}

func (conn *connection) GetRemoteIPAddr() net.IP {
	conn.locker.Lock()
	defer conn.locker.Unlock()
	if conn.remoteAddr == nil {
		return nil
	}
	switch remoteAddr := conn.remoteAddr.(type) {
	case *net.TCPAddr:
		return remoteAddr.IP
	case *net.UDPAddr:
		return remoteAddr.IP
	default:
		panic(fmt.Errorf("unexpected remote addr type %T", remoteAddr))
	}
}

func (conn *connection) GetRemotePort() uint16 {
	conn.locker.Lock()
	defer conn.locker.Unlock()
	if conn.remoteAddr == nil {
		return 0
	}
	switch remoteAddr := conn.remoteAddr.(type) {
	case *net.TCPAddr:
		return uint16(remoteAddr.Port)
	case *net.UDPAddr:
		return uint16(remoteAddr.Port)
	default:
		panic(fmt.Errorf("unexpected remote addr type %T", remoteAddr))
	}
}

func (conn *connection) GetLocalAddr() net.Addr {
	conn.locker.Lock()
	defer conn.locker.Unlock()
	if conn.conn == nil {
		return nil
	}
	return conn.conn.LocalAddr()
}

func (conn *connection) GetLocalIPAddr() net.IP {
	conn.locker.Lock()
	defer conn.locker.Unlock()
	if conn.conn == nil {
		return nil
	}
	localAddr := conn.conn.LocalAddr()
	if localAddr == nil {
		return nil
	}
	switch localAddr := localAddr.(type) {
	case *net.TCPAddr:
		return localAddr.IP
	case *net.UDPAddr:
		return localAddr.IP
	default:
		panic(fmt.Errorf("unexpected remote addr type %T", localAddr))
	}
}

func (conn *connection) GetLocalPort() uint16 {
	conn.locker.Lock()
	defer conn.locker.Unlock()
	if conn.conn == nil {
		return 0
	}
	localAddr := conn.conn.LocalAddr()
	if localAddr == nil {
		return 0
	}
	switch localAddr := localAddr.(type) {
	case *net.TCPAddr:
		return uint16(localAddr.Port)
	case *net.UDPAddr:
		return uint16(localAddr.Port)
	default:
		panic(fmt.Errorf("unexpected local addr type %T", localAddr))
	}
}

func (conn *connection) Close() error {
	conn.cancelFunc()
	conn.waitGroup.Wait()
	return nil
}

func (conn *connection) serve(
	ctx context.Context,
	replyServer proxy_grpc.NetworkProxy_ProxyServer,
) error {
	if conn.cancelFunc != nil {
		panic("serve is not supposed to be called multiple times")
	}
	ctx, cancelFunc := context.WithCancel(ctx)
	conn.cancelFunc = cancelFunc

	go conn.statusSenderLoop(ctx, replyServer)
	err := conn.connect(ctx)
	if err != nil {
		return ErrConnect{Err: err}
	}

	err = conn.loop(ctx, replyServer)
	if err != nil {
		conn.Close()
		return ErrProxy{Err: err}
	}

	return fmt.Errorf("internal error: this line was supposed to be unreachable")
}

func (conn *connection) loop(
	ctx context.Context,
	replyServer proxy_grpc.NetworkProxy_ProxyServer,
) error {
	conn.waitGroup.Add(1)
	go func() {
		defer conn.waitGroup.Done()
		conn.targetReaderLoop(ctx)
	}()

	conn.waitGroup.Add(1)
	go func() {
		defer conn.waitGroup.Done()
		conn.clientReaderLoop(ctx, replyServer)
	}()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case payload := <-conn.receivedFromTarget:
			err := conn.sendToClient(ctx, replyServer, payload)
			if err != nil {
				return fmt.Errorf("unable to forward a payload to the client: %w", err)
			}
			conn.readTargetBufferLocker.Unlock()
		case payload := <-conn.receivedFromClient:
			err := conn.sendToTarget(ctx, payload)
			if err != nil {
				return fmt.Errorf("unable to forward a payload to the target: %w", err)
			}
		}
	}
}

func (conn *connection) sendToClient(
	_ context.Context,
	replyServer proxy_grpc.NetworkProxy_ProxyServer,
	payload []byte,
) error {
	return replyServer.Send(&proxy_grpc.MessageBack{
		MessageBackOneOf: &proxy_grpc.MessageBack_ReceivedPayload{
			ReceivedPayload: payload,
		},
	})
}

func (conn *connection) sendToTarget(
	_ context.Context,
	payload []byte,
) error {
	n, err := conn.conn.Write(payload)
	if err != nil {
		return fmt.Errorf("unable to write to the connection: %w", err)
	}
	if n != len(payload) {
		return fmt.Errorf("incomplete sending happened: %d < %d", n, len(payload))
	}
	return nil
}

func (conn *connection) targetReaderLoop(
	_ context.Context,
) error {
	for {
		n, err := conn.conn.Read(conn.readerTargetBuffer[:])
		if err != nil {
			conn.setStatus(proxy_grpc.ConnectionStatus_BrokenProxyConnection, err)
		}
		conn.readTargetBufferLocker.Lock()
		copy(conn.readTargetBuffer[:n], conn.readerTargetBuffer[:n])
		payload := conn.readTargetBuffer[:n]
		conn.receivedFromTarget <- payload
	}
}

func (conn *connection) clientReaderLoop(
	_ context.Context,
	replyServer proxy_grpc.NetworkProxy_ProxyServer,
) error {
	for {
		msg, err := replyServer.Recv()
		if err != nil {
			return fmt.Errorf("unable to receive a message: %w", err)
		}

		oneOf := msg.GetMessageForwardOneOf()
		payloadMsg, ok := oneOf.(*proxy_grpc.MessageForward_SendPayload)
		if !ok {
			return fmt.Errorf("unexpected a message of type %T, but received %T", payloadMsg, oneOf)
		}

		conn.receivedFromClient <- payloadMsg.SendPayload
	}
}

func (conn *connection) statusSenderLoop(
	ctx context.Context,
	replyServer proxy_grpc.NetworkProxy_ProxyServer,
) {
	for {
		select {
		case <-ctx.Done():
			return
		case s := <-conn.statusChan:
			var errMsg string
			if s.Err != nil {
				errMsg = s.Err.Error()
			}
			statusUpdate := &proxy_grpc.ConnectionStatusUpdate{
				ConnectionID:     uint64(conn.GetID()),
				ConnectionStatus: s.Status,
				ErrorMessage:     ptr(errMsg),
				LocalIPAddress:   conn.GetLocalIPAddr(),
				RemoteIPAddress:  conn.GetRemoteIPAddr(),
			}
			if port := conn.GetRemotePort(); port != 0 {
				statusUpdate.RemotePort = ptr(uint32(port))
			}
			if port := conn.GetLocalPort(); port != 0 {
				statusUpdate.LocalPort = ptr(uint32(port))
			}
			msg := &proxy_grpc.MessageBack{
				MessageBackOneOf: &proxy_grpc.MessageBack_ConnectionUpdate{
					ConnectionUpdate: statusUpdate,
				},
			}
			err := replyServer.Send(msg)
			if err != nil {
				logger.Errorf(ctx, "unable to send status %#+v: %w", msg, err)
				return
			}
		}
	}
}

func (conn *connection) setStatus(
	status proxy_grpc.ConnectionStatus,
	err error,
) {
	conn.statusChan <- connectionStatus{
		Err:    err,
		Status: status,
	}
}

func (conn *connection) connect(
	ctx context.Context,
) error {
	addrs, err := conn.getRemoteAddrs()
	if err != nil {
		conn.setStatus(proxy_grpc.ConnectionStatus_UnableToResolve, err)
		return err
	}

	if len(addrs) == 0 {
		err = fmt.Errorf("received zero IP addresses")
		conn.setStatus(proxy_grpc.ConnectionStatus_UnableToResolve, err)
		return err
	}

	var result *multierror.Error
	for _, addr := range addrs {
		conn.setStatus(proxy_grpc.ConnectionStatus_Connecting, nil)

		conn.setRemoteAddr(addr)
		err := conn.tryConnect(ctx, addr)
		if err == nil {
			conn.setStatus(proxy_grpc.ConnectionStatus_Connected, nil)
			return nil
		}
		result = multierror.Append(result, err)
	}

	return result.ErrorOrNil()
}

func (conn *connection) tryConnect(
	_ context.Context,
	addr net.Addr,
) error {
	var err error
	var c net.Conn
	switch addr := addr.(type) {
	case *net.TCPAddr:
		c, err = net.DialTCP("tcp", nil, addr)
	case *net.UDPAddr:
		c, err = net.DialUDP("udp", nil, addr)
	default:
		return fmt.Errorf("unexpected addr type %T", addr)
	}
	if err != nil {
		return fmt.Errorf("unable to dial %#+v: %w", addr, err)
	}
	conn.setConn(c)
	return nil
}

func (conn *connection) setConn(
	c net.Conn,
) {
	conn.locker.Lock()
	defer conn.locker.Unlock()
	conn.conn = c
	conn.remoteAddr = c.RemoteAddr()
}

func (conn *connection) setRemoteAddr(addr net.Addr) {
	conn.locker.Lock()
	defer conn.locker.Unlock()
	conn.remoteAddr = addr
}

func (conn *connection) getRemoteAddrs() ([]net.Addr, error) {
	addrs, err := net.LookupIP(conn.host)
	if err != nil {
		return nil, ErrResolve{Err: err}
	}

	var result []net.Addr
	for _, addr := range addrs {
		result = append(result, newAddr(conn.proto, addr, conn.port))
	}
	return result, nil
}

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
	case proxy_grpc.NetworkProtocol_UDP:
		return &net.UDPAddr{
			IP:   ipAddr,
			Port: int(port),
		}
	default:
		panic(fmt.Errorf("unexpected proto: %v", proto))
	}
}

func ProtocolToNetwork(proto proxy_grpc.NetworkProtocol) string {
	switch proto {
	case proxy_grpc.NetworkProtocol_TCP:
		return "tcp"
	case proxy_grpc.NetworkProtocol_UDP:
		return "udp"
	default:
		return ""
	}
}
