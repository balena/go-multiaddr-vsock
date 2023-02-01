package mavs

import (
	"encoding/binary"
	"fmt"
	"strconv"

	ma "github.com/multiformats/go-multiaddr"
)

var TranscoderVsock = ma.NewTranscoderFromFunctions(vsockStB, vsockBtS, nil)

func vsockBtS(b []byte) (string, error) {
	contextID := binary.BigEndian.Uint32(b)
	if contextID == 0 {
		return "x", nil
	} else {
		return fmt.Sprintf("%d", contextID), nil
	}
}

func vsockStB(s string) ([]byte, error) {
	b := make([]byte, 4)
	if s == "x" {
		binary.BigEndian.PutUint32(b, 0)
	} else {
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("failed to parse vsock cid: %s", err)
		}
		if i >= 4294967296 {
			return nil, fmt.Errorf("failed to parse vsock cid: %s", "greater than 4294967296")
		}
		binary.BigEndian.PutUint32(b, uint32(i))
	}
	return b, nil
}
