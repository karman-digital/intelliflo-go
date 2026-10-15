package marketingpreferences

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/karman-digital/intelliflo-go/intelliflo/api/credentials"
	marketingpreferencemodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/person/marketingpreferences"
	sharedmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/shared"
	"github.com/karman-digital/intelliflo-go/intelliflo/shared"
)

type marketingPreferenceCredentials struct {
	credentials.Credentials
	statusCode int
}

func (c *marketingPreferenceCredentials) SendRequest(string, string, []byte, ...sharedmodels.GetOptions) (*http.Response, error) {
	return &http.Response{
		StatusCode: c.statusCode,
		Body:       io.NopCloser(bytes.NewBufferString(`{"message":"request failed"}`)),
	}, nil
}

func (c *marketingPreferenceCredentials) Client() *retryablehttp.Client           { return nil }
func (c *marketingPreferenceCredentials) AccessToken() *sharedmodels.AccessToken  { return nil }
func (c *marketingPreferenceCredentials) ApiKey() sharedmodels.APIKey             { return "" }
func (c *marketingPreferenceCredentials) ClientSecret() sharedmodels.ClientSecret { return "" }
func (c *marketingPreferenceCredentials) ClientId() sharedmodels.ClientId         { return "" }

func TestGetMarketingPreferencePreservesResourceNotFound(t *testing.T) {
	service := NewMarketingPreferencesService(&marketingPreferenceCredentials{statusCode: http.StatusNotFound})

	_, err := service.GetMarketingPreference(999999999)
	if !errors.Is(err, shared.ErrResourceNotFound) {
		t.Fatalf("error = %v, want ErrResourceNotFound", err)
	}
	if !strings.Contains(err.Error(), "request failed") {
		t.Fatalf("error = %v, want provider response body", err)
	}
}

func TestPutMarketingPreferencePreservesRequestRejected(t *testing.T) {
	service := NewMarketingPreferencesService(&marketingPreferenceCredentials{statusCode: http.StatusBadRequest})

	_, err := service.PutMarketingPreference(41589479, marketingpreferencemodels.Preferences{})
	if !errors.Is(err, shared.ErrRequestRejected) {
		t.Fatalf("error = %v, want ErrRequestRejected", err)
	}
	if !strings.Contains(err.Error(), "request failed") {
		t.Fatalf("error = %v, want provider response body", err)
	}
}
