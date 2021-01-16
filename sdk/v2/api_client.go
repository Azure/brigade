package sdk

import (
	"github.com/brigadecore/brigade/sdk/v2/authx"
	"github.com/brigadecore/brigade/sdk/v2/core"
	"github.com/brigadecore/brigade/sdk/v2/restmachinery"
)

// APIClient is the general interface for the Brigade API. It does little more
// than expose functions for obtaining more specialized clients for different
// areas of concern, like User management or Project management.
type APIClient interface {
	Authx() authx.APIClient
	Core() core.APIClient
}

type apiClient struct {
	authxClient authx.APIClient
	coreClient  core.APIClient
}

// NewAPIClient returns a Brigade client.
func NewAPIClient(
	apiAddress string,
	apiToken string,
	opts *restmachinery.APIClientOptions,
) APIClient {
	return &apiClient{
		authxClient: authx.NewAPIClient(apiAddress, apiToken, opts),
		coreClient:  core.NewAPIClient(apiAddress, apiToken, opts),
	}
}

func (a *apiClient) Authx() authx.APIClient {
	return a.authxClient
}

func (a *apiClient) Core() core.APIClient {
	return a.coreClient
}
