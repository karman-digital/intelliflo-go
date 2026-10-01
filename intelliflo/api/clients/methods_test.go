package clients

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/karman-digital/intelliflo-go/intelliflo/api/credentials"
	sharedmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/shared"
	"github.com/karman-digital/intelliflo-go/intelliflo/shared"
)

type clientCredentials struct {
	credentials.Credentials
}

func (c *clientCredentials) SendRequest(string, string, []byte, ...sharedmodels.GetOptions) (*http.Response, error) {
	return &http.Response{
		StatusCode: http.StatusNotFound,
		Body:       io.NopCloser(bytes.NewBufferString(`{"message":"not found"}`)),
	}, nil
}

func (c *clientCredentials) Client() *retryablehttp.Client           { return nil }
func (c *clientCredentials) AccessToken() *sharedmodels.AccessToken  { return nil }
func (c *clientCredentials) ApiKey() sharedmodels.APIKey             { return "" }
func (c *clientCredentials) ClientSecret() sharedmodels.ClientSecret { return "" }
func (c *clientCredentials) ClientId() sharedmodels.ClientId         { return "" }

func TestGetClientPreservesResourceNotFound(t *testing.T) {
	service := NewClientService(&clientCredentials{})

	_, err := service.GetClient(999999999)
	if !errors.Is(err, shared.ErrResourceNotFound) {
		t.Fatalf("error = %v, want ErrResourceNotFound", err)
	}
}
