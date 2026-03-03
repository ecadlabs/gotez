package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetRoundFromTenderbakeBlock(t *testing.T) {
	// Helper to build valid tenderbake fitness bytes.
	// Format: <version_len(4)><version(1)>
	//         <level_len(4)><level(4)>
	//         <locked_round_len(4)><locked_round(0 or 4)>
	//         <predecessor_round_len(4)><predecessor_round(4)>
	//         <round_len(4)><round(4)>
	buildFitness := func(level uint32, lockedRound *uint32, predecessorRound uint32, round uint32) []byte {
		b := []byte{
			0, 0, 0, 1, 0x02, // version_len=1, version=2
			0, 0, 0, 4, byte(level >> 24), byte(level >> 16), byte(level >> 8), byte(level), // level
		}
		if lockedRound == nil {
			b = append(b, 0, 0, 0, 0) // locked_round_len=0
		} else {
			lr := *lockedRound
			b = append(b, 0, 0, 0, 4, byte(lr>>24), byte(lr>>16), byte(lr>>8), byte(lr))
		}
		b = append(b, 0, 0, 0, 4, byte(predecessorRound>>24), byte(predecessorRound>>16), byte(predecessorRound>>8), byte(predecessorRound))
		b = append(b, 0, 0, 0, 4, byte(round>>24), byte(round>>16), byte(round>>8), byte(round))
		return b
	}

	uint32Ptr := func(v uint32) *uint32 { return &v }

	t.Run("real mainnet fitness round 0", func(t *testing.T) {
		// From proto_016_PtMumbai operations_test.go (double_baking_evidence),
		// mainnet level 3365599. Fitness version 0x02 (Tenderbake, stable since proto_012).
		fitness := []byte{
			0x00, 0x00, 0x00, 0x01, 0x02,
			0x00, 0x00, 0x00, 0x04, 0x00, 0x33, 0x5a, 0xdf,
			0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x04, 0xff, 0xff, 0xff, 0xff,
			0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00,
		}
		round, err := GetRoundFromTenderbakeBlock(fitness)
		require.NoError(t, err)
		assert.Equal(t, uint32(0), round)
	})

	t.Run("real mainnet fitness round 1", func(t *testing.T) {
		// Mainnet level 12169657 (proto_024), block round 1 (reproposal).
		// Fitness from RPC: ["02","00b9b1b9","","ffffffff","00000001"]
		fitness := []byte{
			0x00, 0x00, 0x00, 0x01, 0x02,
			0x00, 0x00, 0x00, 0x04, 0x00, 0xb9, 0xb1, 0xb9,
			0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x04, 0xff, 0xff, 0xff, 0xff,
			0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x01,
		}
		round, err := GetRoundFromTenderbakeBlock(fitness)
		require.NoError(t, err)
		assert.Equal(t, uint32(1), round)
	})

	t.Run("round extraction without locked round", func(t *testing.T) {
		fitness := buildFitness(100, nil, 0, 5)
		round, err := GetRoundFromTenderbakeBlock(fitness)
		require.NoError(t, err)
		assert.Equal(t, uint32(5), round)
	})

	t.Run("round extraction with locked round", func(t *testing.T) {
		// With locked round present, fitness is 37 bytes instead of 33
		fitness := buildFitness(100, uint32Ptr(2), 0, 7)
		round, err := GetRoundFromTenderbakeBlock(fitness)
		require.NoError(t, err)
		assert.Equal(t, uint32(7), round)
	})

	t.Run("reproposal has different round than original", func(t *testing.T) {
		// This is the actual scenario that caused the bug: a block reproposed
		// at a higher round. The payload_round stays at 0 but the fitness
		// round is 3.
		fitness := buildFitness(100, uint32Ptr(0), 0, 3)
		round, err := GetRoundFromTenderbakeBlock(fitness)
		require.NoError(t, err)
		assert.Equal(t, uint32(3), round, "reproposal round should come from fitness, not payload_round")
	})

	t.Run("rejects truncated fitness", func(t *testing.T) {
		_, err := GetRoundFromTenderbakeBlock([]byte{0, 0, 0, 1, 2})
		assert.Error(t, err)
	})

	t.Run("rejects empty fitness", func(t *testing.T) {
		_, err := GetRoundFromTenderbakeBlock(nil)
		assert.Error(t, err)
	})
}
