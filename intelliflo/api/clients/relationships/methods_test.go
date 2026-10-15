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
	body       string
}

func (c *relationshipCredentials) SendRequest(string, string, []byte, ...sharedmodels.GetOptions) (*http.Response, error) {
	body := c.body
	if body == "" {
		body = `{"message":"request failed"}`
	}
	return &http.Response{
		StatusCode: c.statusCode,
		Body:       io.NopCloser(bytes.NewBufferString(body)),
	}, nil
}

func (c *relationshipCredentials) Client() *retryablehttp.Client           { return nil }
func (c *relationshipCredentials) AccessToken() *sharedmodels.AccessToken  { return nil }
func (c *relationshipCredentials) ApiKey() sharedmodels.APIKey             { return "" }
func (c *relationshipCredentials) ClientSecret() sharedmodels.ClientSecret { return "" }
func (c *relationshipCredentials) ClientId() sharedmodels.ClientId         { return "" }

func TestPostRelationshipAcceptsCreatedResponse(t *testing.T) {
	service := NewRelationshipService(&relationshipCredentials{
		statusCode: http.StatusCreated,
		body:       `{"id":22559522,"relationshipType":{"name":"Other"},"subject":{"id":41589479,"type":"Client","partyType":"Person"},"relation":{"id":41615046,"type":"Client","partyType":"Person"}}`,
	})

	got, err := service.PostRelationship(41589479, relationshipsmodels.Relationship{})
	if err != nil {
		t.Fatalf("PostRelationship returned error: %v", err)
	}
	if got.ID != 22559522 {
		t.Fatalf("relationship ID = %d, want 22559522", got.ID)
	}
}

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
