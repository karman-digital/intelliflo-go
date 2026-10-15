package relationships

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/karman-digital/intelliflo-go/intelliflo/api/credentials"
	relationshipsmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/person/relationships"
	sharedmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/shared"
	"github.com/karman-digital/intelliflo-go/intelliflo/shared"
)

type relationshipCredentials struct {
	credentials.Credentials
	statusCode int
}

func (c *relationshipCredentials) SendRequest(string, string, []byte, ...sharedmodels.GetOptions) (*http.Response, error) {
	return &http.Response{
		StatusCode: c.statusCode,
		Body:       io.NopCloser(bytes.NewBufferString(`{"message":"request failed"}`)),
	}, nil
}

func (c *relationshipCredentials) Client() *retryablehttp.Client           { return nil }
func (c *relationshipCredentials) AccessToken() *sharedmodels.AccessToken  { return nil }
func (c *relationshipCredentials) ApiKey() sharedmodels.APIKey             { return "" }
func (c *relationshipCredentials) ClientSecret() sharedmodels.ClientSecret { return "" }
func (c *relationshipCredentials) ClientId() sharedmodels.ClientId         { return "" }

func TestPostRelationshipPreservesResourceNotFound(t *testing.T) {
	service := NewRelationshipService(&relationshipCredentials{statusCode: http.StatusNotFound})

	_, err := service.PostRelationship(999999999, relationshipsmodels.Relationship{})
	if !errors.Is(err, shared.ErrResourceNotFound) {
		t.Fatalf("error = %v, want ErrResourceNotFound", err)
	}
	if !strings.Contains(err.Error(), "request failed") {
		t.Fatalf("error = %v, want provider response body", err)
	}
}
