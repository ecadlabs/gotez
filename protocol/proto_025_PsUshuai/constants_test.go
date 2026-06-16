package proto_025_PsUshuai

import (
	_ "embed"
	"fmt"
	"testing"

	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:embed constants.bin
var constantsSrc []byte

func TestConstants(t *testing.T) {
	var out Constants
	rest, err := encoding.Decode(constantsSrc, &out, encoding.Dynamic())
	if !assert.NoError(t, err) {
		if err, ok := err.(*encoding.Error); ok {
			fmt.Println(err.Path)
		}
		return
	}
	require.Empty(t, rest)

	require.Equal(t, int8(5), out.CacheStakeInfoCycles)
	require.Equal(t, int8(5), out.CacheSwrrSelectedDistributionCycles)
	require.True(t, out.DALParametric.DynamicLagEnable)
	require.Equal(t, []uint8{1, 2, 3, 4, 5}, out.DALParametric.AttestationLags)
	require.Equal(t, int32(380832), out.DALParametric.SlotSize)
	require.True(t, out.CanonicalRollupAddress.IsNone())
	require.False(t, out.NativeContractsEnable)
	require.False(t, out.SwrrNewBakerLotteryEnable)
	require.False(t, out.Tz5AccountEnable)
	require.Equal(t, uint8(2), out.IssuanceModificationDelay)
	require.Equal(t, uint8(2), out.ConsensusKeyActivationDelay)
	require.Equal(t, uint8(3), out.UnstakeFinalizationDelay)
}
