package libp2pvsock

import (
	manet "github.com/multiformats/go-multiaddr/net"
)

func init() {
	manet.RegisterFromNetAddr(parseVsockNetAddr, "vsock")
	manet.RegisterToNetAddr(parseVsockNetMaddr, "vsock")
}
