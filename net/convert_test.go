package mavsnet

import (
	"testing"

	ma "github.com/multiformats/go-multiaddr"
)

func TestDialArgs(t *testing.T) {
	tests := []struct {
		input           string
		contextID, port uint32
		err             error
	}{
		{
			"/vsock/3/tcp/55555", 3, 55555, nil,
		},
		{
			"/vsock/x/tcp/55555", 0, 55555, nil,
		},
	}

	for _, test := range tests {
		m, _ := ma.NewMultiaddr(test.input)
		contextID, port, err := DialArgs(m)
		if contextID != test.contextID {
			t.Errorf("contextID: expected %d, got %d", test.contextID, contextID)
		}
		if port != test.port {
			t.Errorf("port: expected %d, got %d", test.port, port)
		}
		if err != test.err {
			t.Errorf("err: expected %d, got %d", test.err, err)
		}
	}
}
