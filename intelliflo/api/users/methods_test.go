package users

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/karman-digital/intelliflo-go/intelliflo/api/credentials"
)

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestGetAllUsers(t *testing.T) {
	var queries []string

	client := retryablehttp.NewClient()
	client.RetryMax = 0
	client.Logger = nil
	client.HTTPClient = &http.Client{
		Transport: roundTripperFunc(func(req *http.Request) (*http.Response, error) {
			queries = append(queries, req.URL.RawQuery)
			switch req.URL.RawQuery {
			case "top=500":
				return jsonResponse(`{"items":[{"id":1,"firstName":"Ada","lastName":"Lovelace","status":"Active"}],"count":2,"next_href":"https://api.gb.intelliflo.net/v2/users?skip=500&top=500"}`), nil
			case "skip=500&top=500":
				return jsonResponse(`{"items":[{"id":2,"firstName":"Grace","lastName":"Hopper","status":"Inactive"}],"count":2}`), nil
			default:
				t.Fatalf("unexpected query: %s", req.URL.RawQuery)
				return nil, nil
			}
		}),
	}

	var creds credentials.TenantCredentials
	creds.SetClient(client)
	creds.SetAccessToken("token")
	creds.SetApiKey("api-key")

	service := UserService{Credentials: &creds}

	got, err := service.GetAllUsers()
	if err != nil {
		t.Fatalf("GetAllUsers() error = %v", err)
	}
	if len(got.Items) != 2 {
		t.Fatalf("len(got.Items) = %d, want 2", len(got.Items))
	}
	if got.Items[0].ID != 1 || got.Items[1].ID != 2 {
		t.Fatalf("unexpected item ids = %v", []int{got.Items[0].ID, got.Items[1].ID})
	}
	if len(queries) != 2 {
		t.Fatalf("len(queries) = %d, want 2", len(queries))
	}
}

func TestGetAllUsersRejectsRepeatedCursor(t *testing.T) {
	var calls int
	client := retryablehttp.NewClient()
	client.RetryMax = 0
	client.Logger = nil
	client.HTTPClient = &http.Client{Transport: roundTripperFunc(func(*http.Request) (*http.Response, error) {
		calls++
		if calls > 2 {
			t.Fatalf("GetAllUsers followed a repeated cursor")
		}
		return jsonResponse(`{"items":[],"next_href":"https://api.gb.intelliflo.net/v2/users?skip=500"}`), nil
	})}

	var creds credentials.TenantCredentials
	creds.SetClient(client)
	creds.SetAccessToken("token")
	creds.SetApiKey("api-key")
	service := UserService{Credentials: &creds}

	_, err := service.GetAllUsers()
	if err == nil || !strings.Contains(err.Error(), "repeated next cursor") {
		t.Fatalf("error = %v, want repeated next cursor", err)
	}
}

func jsonResponse(body string) *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(body)),
		Header:     make(http.Header),
	}
}
