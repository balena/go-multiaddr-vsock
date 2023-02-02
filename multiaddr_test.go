package mavs

import (
	"testing"

	ma "github.com/multiformats/go-multiaddr"
)

func TestConstructFails(t *testing.T) {
	cases := []string{
		"/vsock",
		"/vsock/abcd",
		"/vsock/3/xtcp/55555/unix",
	}

	for _, a := range cases {
		if _, err := ma.NewMultiaddr(a); err == nil {
			t.Errorf("should have failed: %s - %s", a, err)
		}
	}
}

func TestConstructSucceeds(t *testing.T) {
	cases := []string{
		"/vsock/3",
		"/vsock/3/tcp/55555",
		"/vsock/3/udp/55555",
		"/vsock/x/xtcp/55555", // gets cid from /dev/vsock
	}

	for _, a := range cases {
		if _, err := ma.NewMultiaddr(a); err != nil {
			t.Errorf("should have succeeded: %s -- %s", a, err)
		}
	}

}
