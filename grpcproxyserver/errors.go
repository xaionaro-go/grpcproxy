package grpcproxyserver

import (
	"fmt"
)

// ErrResolve means an error occurred while resolving the hostname.
type ErrResolve struct {
	Err error
}

func (e ErrResolve) Error() string {
	return fmt.Sprintf("unable to resolve: %v", e.Err)
}

func (e ErrResolve) Unwrap() error {
	return e.Err
}

// ErrConnect means an error occurred while connecting to the destination.
type ErrConnect struct {
	Err error
}

func (e ErrConnect) Error() string {
	return fmt.Sprintf("unable to connect: %v", e.Err)
}

func (e ErrConnect) Unwrap() error {
	return e.Err
}

// ErrConnect means an error occurred while forwarding traffic.
type ErrProxy struct {
	Err error
}

func (e ErrProxy) Error() string {
	return fmt.Sprintf("unable to proxy/forward traffic: %v", e.Err)
}

func (e ErrProxy) Unwrap() error {
	return e.Err
}
