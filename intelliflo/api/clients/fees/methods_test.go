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
	method    string
	path      string
	responses []string
	options   []sharedmodels.GetOptions
}

func (c *feeCredentials) SendRequest(method, path string, _ []byte, options ...sharedmodels.GetOptions) (*http.Response, error) {
	c.method = method
	c.path = path
	c.options = append(c.options, options...)
	response := "{\"id\":14843314,\"plans\":[{\"id\":77673534}]}"
	if len(c.responses) > 0 {
		response = c.responses[0]
		c.responses = c.responses[1:]
	}
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(response)),
	}, nil
}

func TestGetFeesGetsEveryPage(t *testing.T) {
	creds := &feeCredentials{responses: []string{
		`{"items":[{"id":13719455}],"count":2,"next_href":"https://api.gb.intelliflo.net/v2/clients/38021306/fees?skip=1&top=1"}`,
		`{"items":[{"id":13719456}],"count":2}`,
	}}
	service := NewFeeService(creds)

	fees, err := service.GetFees(38021306)

	require.NoError(t, err)
	assert.Equal(t, []int{13719455, 13719456}, []int{fees.Items[0].ID, fees.Items[1].ID})
	require.Len(t, creds.options, 2)
	assert.Equal(t, 500, creds.options[0].Top)
	assert.Equal(t, 1, creds.options[1].Skip)
}

func TestGetFeesAcceptsNumericPaidBy(t *testing.T) {
	creds := &feeCredentials{responses: []string{
		`{"items":[{"id":14625740,"paymentType":{"paidBy":0}}],"count":1}`,
	}}
	service := NewFeeService(creds)

	fees, err := service.GetFees(38015539)

	require.NoError(t, err)
	require.Len(t, fees.Items, 1)
	assert.Equal(t, "0", string(fees.Items[0].PaymentType.PaidBy))
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
