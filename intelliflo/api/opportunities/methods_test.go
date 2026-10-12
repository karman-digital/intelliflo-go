package opportunities

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/karman-digital/intelliflo-go/intelliflo/api/credentials"
	opportunitiesmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/opportunities"
	sharedmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/shared"
	"github.com/karman-digital/intelliflo-go/intelliflo/shared"
)

type opportunityCredentials struct {
	credentials.Credentials
}

func (c *opportunityCredentials) SendRequest(string, string, []byte, ...sharedmodels.GetOptions) (*http.Response, error) {
	return &http.Response{
		StatusCode: http.StatusNotFound,
		Body:       io.NopCloser(bytes.NewBufferString(`{"message":"not found"}`)),
	}, nil
}

func (c *opportunityCredentials) Client() *retryablehttp.Client           { return nil }
func (c *opportunityCredentials) AccessToken() *sharedmodels.AccessToken  { return nil }
func (c *opportunityCredentials) ApiKey() sharedmodels.APIKey             { return "" }
func (c *opportunityCredentials) ClientSecret() sharedmodels.ClientSecret { return "" }
func (c *opportunityCredentials) ClientId() sharedmodels.ClientId         { return "" }

func TestCreateClientOpportunityPreservesResourceNotFound(t *testing.T) {
	service := NewOpportunityService(&opportunityCredentials{})

	_, err := service.CreateClientOpportunity(999999999, opportunitiesmodels.OpportunityCreateRequest{})
	if !errors.Is(err, shared.ErrResourceNotFound) {
		t.Fatalf("error = %v, want ErrResourceNotFound", err)
	}
}
