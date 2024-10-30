package grpcproxyserver

import (
	"fmt"
)

type ErrResolve struct {
	Err error
}

func (e ErrResolve) Error() string {
	return fmt.Sprintf("unable to resolve: %v", e.Err)
}

func (e ErrResolve) Unwrap() error {
	return e.Err
}

type ErrConnect struct {
	Err error
}

func (e ErrConnect) Error() string {
	return fmt.Sprintf("unable to connect: %v", e.Err)
}

func (e ErrConnect) Unwrap() error {
	return e.Err
}

type ErrProxy struct {
	Err error
}

func (e ErrProxy) Error() string {
	return fmt.Sprintf("unable to proxy/forward traffic: %v", e.Err)
}

func (e ErrProxy) Unwrap() error {
	return e.Err
}
