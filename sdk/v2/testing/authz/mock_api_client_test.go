package authz

import (
	"testing"

	"github.com/brigadecore/brigade/sdk/v2/authz"
	"github.com/stretchr/testify/require"
)

func TestMockAPIClient(t *testing.T) {
	require.Implements(t, (*authz.APIClient)(nil), &MockAPIClient{})
}
