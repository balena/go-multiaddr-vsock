package mavs

import (
	ma "github.com/multiformats/go-multiaddr"
)

const (
	P_VSOCK = 40
	P_XTCP  = 7777972
)

var (
	protoVSOCK = ma.Protocol{
		Name:       "vsock",
		Code:       P_VSOCK,
		VCode:      ma.CodeToVarint(P_VSOCK),
		Size:       ma.LengthPrefixedVarSize,
		Transcoder: TranscoderVsock,
	}
	protoXTCP = ma.Protocol{
		Name:       "xtcp",
		Code:       P_XTCP,
		VCode:      ma.CodeToVarint(P_XTCP),
		Size:       32,
		Path:       false,
		Transcoder: TranscoderXport,
	}
)

func init() {
	for _, p := range []ma.Protocol{
		protoVSOCK,
		protoXTCP,
	} {
		if err := ma.AddProtocol(p); err != nil {
			panic(err)
		}
	}
}
