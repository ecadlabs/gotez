package proto_023_PtSeouLo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListSignRequests(t *testing.T) {
	reqs := ListSignRequests()
	expect := []string{
		"attestation",
		"attestation_with_dal",
		"block",
		"generic",
		"pack",
		"preattestation",
	}
	require.Equal(t, expect, reqs)
}
