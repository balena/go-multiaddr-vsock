package mavsnet

import (
	"encoding/binary"
	"fmt"
	"net"

	"github.com/mdlayher/vsock"
	ma "github.com/multiformats/go-multiaddr"
)

var errIncorrectNetAddr = fmt.Errorf("incorrect network addr conversion")

func parseVsockNetAddr(a net.Addr) (ma.Multiaddr, error) {
	ac, ok := a.(*vsock.Addr)
	if !ok {
		return nil, errIncorrectNetAddr
	}

	return ma.NewMultiaddr(fmt.Sprintf("/vsock/%d/tcp/%d", ac.ContextID, ac.Port))
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
			case P_VSOCK:
				network = "vsock"
				contextID = binary.BigEndian.Uint32(c.Bytes())
				return true
			}
		case "vsock":
			switch c.Protocol().Code {
			case ma.P_TCP:
				network = "tcp"
			default:
				err = fmt.Errorf("%s has unsupported tx", m)
				return false
			}
			port = binary.BigEndian.Uint32(c.Bytes())
		}
		// Done.
		return false
	})
	return
}
