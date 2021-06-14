package core

import (
	"testing"

	"github.com/brigadecore/brigade/sdk/v2/core"
	"github.com/stretchr/testify/require"
)

func TestMockProjectRoleAssignmentsClient(t *testing.T) {
	require.Implements(
		t,
		(*core.ProjectRoleAssignmentsClient)(nil),
		&MockProjectRoleAssignmentsClient{},
	)
}