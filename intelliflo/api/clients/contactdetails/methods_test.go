package contactdetails

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/karman-digital/intelliflo-go/intelliflo/api/credentials"
	contactdetailmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/person/contactdetails"
	sharedmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/shared"
	"github.com/karman-digital/intelliflo-go/intelliflo/shared"
)

type contactDetailCredentials struct {
	credentials.Credentials
	statusCode int
}

func (c *contactDetailCredentials) SendRequest(string, string, []byte, ...sharedmodels.GetOptions) (*http.Response, error) {
	return &http.Response{
		StatusCode: c.statusCode,
		Body:       io.NopCloser(bytes.NewBufferString(`{"message":"request failed"}`)),
	}, nil
}

func (c *contactDetailCredentials) Client() *retryablehttp.Client           { return nil }
func (c *contactDetailCredentials) AccessToken() *sharedmodels.AccessToken  { return nil }
func (c *contactDetailCredentials) ApiKey() sharedmodels.APIKey             { return "" }
func (c *contactDetailCredentials) ClientSecret() sharedmodels.ClientSecret { return "" }
func (c *contactDetailCredentials) ClientId() sharedmodels.ClientId         { return "" }

func TestGetContactDetailsPreservesResourceNotFound(t *testing.T) {
	service := NewContactDetailsService(&contactDetailCredentials{statusCode: http.StatusNotFound})

	_, err := service.GetContactDetails(999999999)
	if !errors.Is(err, shared.ErrResourceNotFound) {
		t.Fatalf("error = %v, want ErrResourceNotFound", err)
	}
}

func TestPostContactDetailPreservesRequestRejected(t *testing.T) {
	service := NewContactDetailsService(&contactDetailCredentials{statusCode: http.StatusBadRequest})

	_, err := service.PostContactDetail(123, contactdetailmodels.ContactDetail{})
	if !errors.Is(err, shared.ErrRequestRejected) {
		t.Fatalf("error = %v, want ErrRequestRejected", err)
	}
}

func TestPutContactDetailPreservesRequestRejected(t *testing.T) {
	service := NewContactDetailsService(&contactDetailCredentials{statusCode: http.StatusBadRequest})

	_, err := service.PutContactDetail(123, 456, contactdetailmodels.ContactDetail{})
	if !errors.Is(err, shared.ErrRequestRejected) {
		t.Fatalf("error = %v, want ErrRequestRejected", err)
	}
}
