package addresses

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/karman-digital/intelliflo-go/intelliflo/api/credentials"
	addressmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/person/address"
	sharedmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/shared"
	"github.com/karman-digital/intelliflo-go/intelliflo/shared"
)

type addressCredentials struct {
	credentials.Credentials
}

func (c *addressCredentials) SendRequest(string, string, []byte, ...sharedmodels.GetOptions) (*http.Response, error) {
	return &http.Response{
		StatusCode: http.StatusNotFound,
		Body:       io.NopCloser(bytes.NewBufferString(`{"message":"not found"}`)),
	}, nil
}

func (c *addressCredentials) Client() *retryablehttp.Client           { return nil }
func (c *addressCredentials) AccessToken() *sharedmodels.AccessToken  { return nil }
func (c *addressCredentials) ApiKey() sharedmodels.APIKey             { return "" }
func (c *addressCredentials) ClientSecret() sharedmodels.ClientSecret { return "" }
func (c *addressCredentials) ClientId() sharedmodels.ClientId         { return "" }

func TestGetAddressesPreservesResourceNotFound(t *testing.T) {
	service := NewAddressService(&addressCredentials{})

	_, err := service.GetAddresses(999999999)
	if !errors.Is(err, shared.ErrResourceNotFound) {
		t.Fatalf("error = %v, want ErrResourceNotFound", err)
	}
}

func TestPostAddressPreservesResourceNotFound(t *testing.T) {
	service := NewAddressService(&addressCredentials{})

	_, err := service.PostAddress(999999999, addressmodels.Residence{})
	if !errors.Is(err, shared.ErrResourceNotFound) {
		t.Fatalf("error = %v, want ErrResourceNotFound", err)
	}
}

func TestPutAddressPreservesResourceNotFound(t *testing.T) {
	service := NewAddressService(&addressCredentials{})

	_, err := service.PutAddress(999999999, 888888888, addressmodels.Residence{})
	if !errors.Is(err, shared.ErrResourceNotFound) {
		t.Fatalf("error = %v, want ErrResourceNotFound", err)
	}
}
