package grpcproxyserver

func ptr[T any](in T) *T {
	return &in
}
