package proto_025_PsUshuai

import (
	"bytes"
	"testing"

	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/stretchr/testify/require"
)

func TestDALEntrapmentEvidenceLagIndexEncodingOrder(t *testing.T) {
	branch := new(tz.BlockHash)
	payloadHash := new(tz.BlockPayloadHash)
	inlined := InlinedConsensusOperation{
		Branch: branch,
		Operations: &Attestation{
			BlockPayloadHash: payloadHash,
		},
		Signature: make(tz.AnySignature, tz.GenericSignatureBytesLen),
	}
	ev := DALEntrapmentEvidence{
		Attestation:   inlined,
		ConsensusSlot: 0x0102,
		SlotIndex:     0x03,
		LagIndex:      tz.Some(uint8(0x04)),
	}

	var inlinedBuf bytes.Buffer
	require.NoError(t, encoding.Encode(&inlinedBuf, &inlined))

	var evBuf bytes.Buffer
	require.NoError(t, encoding.Encode(&evBuf, &ev))

	data := evBuf.Bytes()
	offset := 4 + inlinedBuf.Len()
	require.GreaterOrEqual(t, len(data), offset+5)
	require.Equal(t, []byte{0x01, 0x02, 0x03, 0xff, 0x04}, data[offset:offset+5])
}
