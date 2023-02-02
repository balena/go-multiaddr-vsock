package mavsnet

import (
	"encoding/binary"
	"fmt"
	"net"

	mavs "github.com/balena/go-multiaddr-vsock"
	"github.com/balena/go-vsock"
	ma "github.com/multiformats/go-multiaddr"
)

var errIncorrectNetAddr = fmt.Errorf("incorrect network addr conversion")

func parseVsockNetAddr(a net.Addr) (ma.Multiaddr, error) {
	ac, ok := a.(*vsock.Addr)
	if !ok {
		return nil, errIncorrectNetAddr
	}

	return ma.NewMultiaddr(fmt.Sprintf("/vsock/%d/xtcp/%d", ac.ContextID, ac.Port))
}

func parseVsockNetMaddr(maddr ma.Multiaddr) (net.Addr, error) {
	contextID, port, err := DialArgs(maddr)
	if err != nil {
		return nil, err
	}

	return &vsock.Addr{
		ContextID: contextID,
		Port:      port,
	}, nil
}

// DialArgs is a convenience function that returns cid and port as
// expected by vsock.Dial.
func DialArgs(m ma.Multiaddr) (contextID, port uint32, err error) {
	var network string

	ma.ForEach(m, func(c ma.Component) bool {
		switch network {
		case "":
			switch c.Protocol().Code {
			case mavs.P_VSOCK:
				network = "vsock"
				contextID = binary.BigEndian.Uint32(c.RawValue())
				return true
			}
		case "vsock":
			switch c.Protocol().Code {
			case mavs.P_XTCP:
				network = "xtcp"
			default:
				err = fmt.Errorf("%s has unsupported tx", m)
				return false
			}
			port = binary.BigEndian.Uint32(c.RawValue())
		}
		// Done.
		return false
	})
	return
}
