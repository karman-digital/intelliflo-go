package fees

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/karman-digital/intelliflo-go/intelliflo/api/credentials"
	sharedmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/shared"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type feeCredentials struct {
	credentials.Credentials
	method string
	path   string
}

func (c *feeCredentials) SendRequest(method, path string, _ []byte, _ ...sharedmodels.GetOptions) (*http.Response, error) {
	c.method = method
	c.path = path
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString("{\"id\":14843314,\"plans\":[{\"id\":77673534}]}")),
	}, nil
}

func (c *feeCredentials) Client() *retryablehttp.Client {
	return nil
}

func (c *feeCredentials) AccessToken() *sharedmodels.AccessToken {
	return nil
}

func (c *feeCredentials) ApiKey() sharedmodels.APIKey {
	return ""
}

func (c *feeCredentials) ClientSecret() sharedmodels.ClientSecret {
	return ""
}

func (c *feeCredentials) ClientId() sharedmodels.ClientId {
	return ""
}

func TestGetFee(t *testing.T) {
	creds := &feeCredentials{}
	service := NewFeeService(creds)

	fee, err := service.GetFee(38021306, 14843314)

	require.NoError(t, err)
	assert.Equal(t, http.MethodGet, creds.method)
	assert.Equal(t, "clients/38021306/fees/14843314", creds.path)
	assert.Equal(t, 14843314, fee.ID)
	require.Len(t, fee.Plans, 1)
	assert.Equal(t, 77673534, fee.Plans[0].ID)
}
