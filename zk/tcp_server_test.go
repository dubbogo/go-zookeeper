package zk

import (
	"errors"
	"net"
	"testing"
)

func WithListenServer(t *testing.T, test func(server string)) {
	t.Helper()
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("Failed to start listen server: %v", err)
	}
	t.Cleanup(func() { _ = l.Close() })
	server := l.Addr().String()

	go func() {
		conn, err := l.Accept()
		if err != nil {
			if !errors.Is(err, net.ErrClosed) {
				t.Logf("Failed to accept connection: %v", err)
			}
			return
		}
		_ = conn.Close()
	}()

	test(server)
}
