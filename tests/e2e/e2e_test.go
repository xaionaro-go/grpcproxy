//go:build e2e_tests
// +build e2e_tests

package e2e

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"sync"
	"testing"

	"github.com/facebookincubator/go-belt/tool/logger"
	xlogrus "github.com/facebookincubator/go-belt/tool/logger/implementation/logrus"
	"github.com/facebookincubator/go-belt/tool/logger/types"
	"github.com/stretchr/testify/require"
	"github.com/xaionaro-go/grpcproxy/grpchttpproxy"
	"github.com/xaionaro-go/grpcproxy/grpcproxyserver"
	"github.com/xaionaro-go/grpcproxy/protobuf/go/proxy_grpc"
	"google.golang.org/grpc"
)

func TestE2E(t *testing.T) {
	var wg sync.WaitGroup

	ctx := logger.CtxWithLogger(context.Background(), xlogrus.Default().WithLevel(logger.LevelTrace))
	logger.Default = func() types.Logger {
		return logger.FromCtx(ctx)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /somePath/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "OK\n")
	})

	finalEndpointListener, err := net.ListenTCP("tcp", &net.TCPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 0,
	})
	require.NoError(t, err)

	wg.Add(1)
	go func() {
		defer wg.Done()
		logger.Debugf(ctx, "the final endpoint server start")
		defer logger.Debugf(ctx, "the final endpoint server ended")

		logger.Infof(ctx, "started the final endpoint server at '%s'", finalEndpointListener.Addr())
		err := http.Serve(finalEndpointListener, mux)
		require.NoError(t, err)
	}()

	grpcServerListener, err := net.ListenTCP("tcp", &net.TCPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 0,
	})
	require.NoError(t, err)

	wg.Add(1)
	go func() {
		defer wg.Done()
		logger.Debugf(ctx, "the gRPC server start")
		defer logger.Debugf(ctx, "the gRPC server ended")

		proxyServer := grpcproxyserver.New()
		grpcServer := grpc.NewServer()
		proxy_grpc.RegisterNetworkProxyServer(grpcServer, proxyServer)
		logger.Infof(ctx, "started the gRPC server at '%s'", grpcServerListener.Addr())
		err = grpcServer.Serve(grpcServerListener)
	}()

	conn, err := grpc.NewClient(grpcServerListener.Addr().String(), grpc.WithInsecure())
	require.NoError(t, err)

	proxyClient := proxy_grpc.NewNetworkProxyClient(conn)

	httpClient := &http.Client{
		Transport: &http.Transport{
			DialContext: grpchttpproxy.NewDialer(proxyClient).DialContext,
		},
	}

	u := &url.URL{
		Scheme: "http",
		Host:   finalEndpointListener.Addr().String(),
		Path:   "/somePath",
	}

	resp, err := httpClient.Get(u.String())
	require.NoError(t, err)
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.Equal(t, "OK\n", string(body))
}
