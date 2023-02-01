package libp2pvsock

import (
	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("vsock-tpt")

/*
type vsockListener struct {
	manet.Listener
	sec int
}

func (ll *vsockListener) Accept() (manet.Conn, error) {
	c, err := ll.Listener.Accept()
	if err != nil {
		return nil, err
	}
	return c, nil
}

type Option func(*TcpTransport) error

func DisableReuseport() Option {
	return func(tr *TcpTransport) error {
		tr.disableReuseport = true
		return nil
	}
}

func WithConnectionTimeout(d time.Duration) Option {
	return func(tr *TcpTransport) error {
		tr.connectTimeout = d
		return nil
	}
}

func WithMetrics() Option {
	return func(tr *TcpTransport) error {
		tr.enableMetrics = true
		return nil
	}
}

// TcpTransport is the TCP transport.
type TcpTransport struct {
	// Connection upgrader for upgrading insecure stream connections to
	// secure multiplex connections.
	upgrader transport.Upgrader

	disableReuseport bool // Explicitly disable reuseport.
	enableMetrics    bool

	// TCP connect timeout
	connectTimeout time.Duration

	rcmgr network.ResourceManager

	reuse reuseport.Transport
}

var _ transport.Transport = &TcpTransport{}

// NewTCPTransport creates a tcp transport object that tracks dialers and listeners
// created. It represents an entire TCP stack (though it might not necessarily be).
func NewTCPTransport(upgrader transport.Upgrader, rcmgr network.ResourceManager, opts ...Option) (*TcpTransport, error) {
	if rcmgr == nil {
		rcmgr = &network.NullResourceManager{}
	}
	tr := &TcpTransport{
		upgrader:       upgrader,
		connectTimeout: defaultConnectTimeout, // can be set by using the WithConnectionTimeout option
		rcmgr:          rcmgr,
	}
	for _, o := range opts {
		if err := o(tr); err != nil {
			return nil, err
		}
	}
	return tr, nil
}

var dialMatcher = mafmt.And(mafmt.IP, mafmt.Base(ma.P_TCP))

// CanDial returns true if this transport believes it can dial the given
// multiaddr.
func (t *TcpTransport) CanDial(addr ma.Multiaddr) bool {
	return dialMatcher.Matches(addr)
}

func (t *TcpTransport) maDial(ctx context.Context, raddr ma.Multiaddr) (manet.Conn, error) {
	// Apply the deadline iff applicable
	if t.connectTimeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, t.connectTimeout)
		defer cancel()
	}

	if t.UseReuseport() {
		return t.reuse.DialContext(ctx, raddr)
	}
	var d manet.Dialer
	return d.DialContext(ctx, raddr)
}

// Dial dials the peer at the remote address.
func (t *TcpTransport) Dial(ctx context.Context, raddr ma.Multiaddr, p peer.ID) (transport.CapableConn, error) {
	connScope, err := t.rcmgr.OpenConnection(network.DirOutbound, true, raddr)
	if err != nil {
		log.Debugw("resource manager blocked outgoing connection", "peer", p, "addr", raddr, "error", err)
		return nil, err
	}
	if err := connScope.SetPeer(p); err != nil {
		log.Debugw("resource manager blocked outgoing connection for peer", "peer", p, "addr", raddr, "error", err)
		connScope.Done()
		return nil, err
	}
	conn, err := t.maDial(ctx, raddr)
	if err != nil {
		connScope.Done()
		return nil, err
	}
	// Set linger to 0 so we never get stuck in the TIME-WAIT state. When
	// linger is 0, connections are _reset_ instead of closed with a FIN.
	// This means we can immediately reuse the 5-tuple and reconnect.
	tryLinger(conn, 0)
	tryKeepAlive(conn, true)
	c := conn
	if t.enableMetrics {
		var err error
		c, err = newTracingConn(conn, true)
		if err != nil {
			connScope.Done()
			return nil, err
		}
	}
	direction := network.DirOutbound
	if ok, isClient, _ := network.GetSimultaneousConnect(ctx); ok && !isClient {
		direction = network.DirInbound
	}
	return t.upgrader.Upgrade(ctx, t, c, direction, p, connScope)
}

// UseReuseport returns true if reuseport is enabled and available.
func (t *TcpTransport) UseReuseport() bool {
	return !t.disableReuseport && ReuseportIsAvailable()
}

func (t *TcpTransport) maListen(laddr ma.Multiaddr) (manet.Listener, error) {
	if t.UseReuseport() {
		return t.reuse.Listen(laddr)
	}
	return manet.Listen(laddr)
}

// Listen listens on the given multiaddr.
func (t *TcpTransport) Listen(laddr ma.Multiaddr) (transport.Listener, error) {
	list, err := t.maListen(laddr)
	if err != nil {
		return nil, err
	}
	if t.enableMetrics {
		list = newTracingListener(&vsockListener{list, 0})
	}
	return t.upgrader.UpgradeListener(t, list), nil
}

// Protocols returns the list of terminal protocols this transport can dial.
func (t *TcpTransport) Protocols() []int {
	return []int{ma.P_TCP}
}

// Proxy always returns false for the TCP transport.
func (t *TcpTransport) Proxy() bool {
	return false
}

func (t *TcpTransport) String() string {
	return "TCP"
}
*/
