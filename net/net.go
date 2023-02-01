package mavsnet

import (
	"github.com/mdlayher/vsock"
	ma "github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"
)

// Dial connects to a remote address. It uses an underlying vsock.Conn,
// then wraps it in a manet.Conn object (with local and remote Multiaddrs).
func Dial(remote ma.Multiaddr) (manet.Conn, error) {
	contextID, port, err := DialArgs(remote)
	if err != nil {
		return nil, err
	}

	vs, err := vsock.Dial(contextID, port, nil)
	if err != nil {
		return nil, err
	}

	return manet.WrapNetConn(vs)
}

// Listen announces on the local network address laddr.
func Listen(laddr ma.Multiaddr) (manet.Listener, error) {

	// get the net.Listen friendly arguments from the remote addr
	contextID, port, err := DialArgs(laddr)
	if err != nil {
		return nil, err
	}
	if contextID == 0 {
		contextID, err = vsock.ContextID()
		if err != nil {
			return nil, err
		}
	}

	nl, err := vsock.ListenContextID(contextID, port, nil)
	if err != nil {
		return nil, err
	}

	return manet.WrapNetListener(nl)
}
