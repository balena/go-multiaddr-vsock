package mavs

import (
	ma "github.com/multiformats/go-multiaddr"
)

const (
	P_VSOCK = 40
)

var (
	protoVSOCK = ma.Protocol{
		Name:       "vsock",
		Code:       P_VSOCK,
		VCode:      ma.CodeToVarint(P_VSOCK),
		Size:       ma.LengthPrefixedVarSize,
		Transcoder: TranscoderVsock,
	}
)

func init() {
	if err := ma.AddProtocol(protoVSOCK); err != nil {
		panic(err)
	}
}
