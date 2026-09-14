package activities

import (
	"bytes"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/karman-digital/intelliflo-go/intelliflo/api/credentials"
)

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestGetCategoriesUsesActivitiesCategoriesPath(t *testing.T) {
	client := retryablehttp.NewClient()
	client.RetryMax = 0
	client.Logger = nil
	client.HTTPClient = &http.Client{Transport: roundTripperFunc(func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("method = %s, want GET", req.Method)
		}
		if req.URL.Path != "/v2/activities/categories" {
			t.Fatalf("path = %s, want /v2/activities/categories", req.URL.Path)
		}
		return activityJSONResponse(http.StatusOK, `{"items":[],"count":0}`), nil
	})}

	service := activityServiceWithClient(client)
	if _, err := service.GetCategories(); err != nil {
		t.Fatalf("GetCategories() error = %v", err)
	}
}

func TestGetTypesUsesActivitiesTypesPath(t *testing.T) {
	client := retryablehttp.NewClient()
	client.RetryMax = 0
	client.Logger = nil
	client.HTTPClient = &http.Client{Transport: roundTripperFunc(func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("method = %s, want GET", req.Method)
		}
		if req.URL.Path != "/v2/activities/types" {
			t.Fatalf("path = %s, want /v2/activities/types", req.URL.Path)
		}
		return activityJSONResponse(http.StatusOK, `{"items":[],"count":0}`), nil
	})}

	service := activityServiceWithClient(client)
	if _, err := service.GetTypes(); err != nil {
		t.Fatalf("GetTypes() error = %v", err)
	}
}

func TestGetAllCategoriesPaginates(t *testing.T) {
	var queries []string
	client := retryablehttp.NewClient()
	client.RetryMax = 0
	client.Logger = nil
	client.HTTPClient = &http.Client{Transport: roundTripperFunc(func(req *http.Request) (*http.Response, error) {
		queries = append(queries, req.URL.RawQuery)
		switch req.URL.RawQuery {
		case "top=500":
			return activityJSONResponse(http.StatusOK, `{"href":"first","items":[{"id":2,"name":"Review"}],"count":2,"next_href":"https://api.gb.intelliflo.net/v2/activities/categories?skip=500&top=500"}`), nil
		case "skip=500&top=500":
			return activityJSONResponse(http.StatusOK, `{"items":[{"id":1,"name":"Call"}],"count":2}`), nil
		default:
			t.Fatalf("unexpected query: %s", req.URL.RawQuery)
			return nil, nil
		}
	})}

	service := activityServiceWithClient(client)
	got, err := service.GetAllCategories()
	if err != nil {
		t.Fatalf("GetAllCategories() error = %v", err)
	}
	if ids := []int{got.Items[0].ID, got.Items[1].ID}; !reflect.DeepEqual(ids, []int{2, 1}) {
		t.Fatalf("item ids = %v, want [2 1]", ids)
	}
	if !reflect.DeepEqual(queries, []string{"top=500", "skip=500&top=500"}) {
		t.Fatalf("queries = %v", queries)
	}
}

func TestGetAllTypesPaginates(t *testing.T) {
	var calls int
	client := retryablehttp.NewClient()
	client.RetryMax = 0
	client.Logger = nil
	client.HTTPClient = &http.Client{Transport: roundTripperFunc(func(req *http.Request) (*http.Response, error) {
		calls++
		if req.URL.Path != "/v2/activities/types" {
			t.Fatalf("path = %s", req.URL.Path)
		}
		if calls == 1 {
			return activityJSONResponse(http.StatusOK, `{"items":[{"id":10,"name":"Call"}],"count":2,"next_href":"https://api.gb.intelliflo.net/v2/activities/types?skip=500"}`), nil
		}
		return activityJSONResponse(http.StatusOK, `{"items":[{"id":20,"name":"Meeting"}],"count":2}`), nil
	})}

	service := activityServiceWithClient(client)
	got, err := service.GetAllTypes()
	if err != nil {
		t.Fatalf("GetAllTypes() error = %v", err)
	}
	if ids := []int{got.Items[0].ID, got.Items[1].ID}; !reflect.DeepEqual(ids, []int{10, 20}) {
		t.Fatalf("item ids = %v, want [10 20]", ids)
	}
}

func TestGetPrioritiesUsesActivitiesPrioritiesPath(t *testing.T) {
	client := retryablehttp.NewClient()
	client.RetryMax = 0
	client.Logger = nil
	client.HTTPClient = &http.Client{Transport: roundTripperFunc(func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet || req.URL.Path != "/v2/activities/priorities" {
			t.Fatalf("request = %s %s", req.Method, req.URL.Path)
		}
		return activityJSONResponse(http.StatusOK, `{"items":[{"id":7,"name":"High"}],"count":1}`), nil
	})}

	service := activityServiceWithClient(client)
	got, err := service.GetPriorities()
	if err != nil {
		t.Fatalf("GetPriorities() error = %v", err)
	}
	if len(got.Items) != 1 || got.Items[0].ID != 7 || got.Items[0].Name != "High" {
		t.Fatalf("items = %#v", got.Items)
	}
}

func TestGetAllPrioritiesPaginates(t *testing.T) {
	var calls int
	client := retryablehttp.NewClient()
	client.RetryMax = 0
	client.Logger = nil
	client.HTTPClient = &http.Client{Transport: roundTripperFunc(func(req *http.Request) (*http.Response, error) {
		calls++
		if calls == 1 {
			return activityJSONResponse(http.StatusOK, `{"items":[{"id":1,"name":"Low"}],"count":2,"next_href":"https://api.gb.intelliflo.net/v2/activities/priorities?skip=500"}`), nil
		}
		return activityJSONResponse(http.StatusOK, `{"items":[{"id":2,"name":"High"}],"count":2}`), nil
	})}

	service := activityServiceWithClient(client)
	got, err := service.GetAllPriorities()
	if err != nil {
		t.Fatalf("GetAllPriorities() error = %v", err)
	}
	if len(got.Items) != 2 || got.Items[1].ID != 2 {
		t.Fatalf("items = %#v", got.Items)
	}
}

func TestGetOutcomesUsesActivitiesOutcomesPath(t *testing.T) {
	client := retryablehttp.NewClient()
	client.RetryMax = 0
	client.Logger = nil
	client.HTTPClient = &http.Client{Transport: roundTripperFunc(func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet || req.URL.Path != "/v2/activities/outcomes" {
			t.Fatalf("request = %s %s", req.Method, req.URL.Path)
		}
		return activityJSONResponse(http.StatusOK, `{"items":[{"id":31,"name":"Completed","activityType":{"id":10}}],"count":1}`), nil
	})}

	service := activityServiceWithClient(client)
	got, err := service.GetOutcomes()
	if err != nil {
		t.Fatalf("GetOutcomes() error = %v", err)
	}
	if len(got.Items) != 1 || got.Items[0].ID != 31 || got.Items[0].ActivityType.ID != 10 {
		t.Fatalf("items = %#v", got.Items)
	}
}

func TestGetAllOutcomesPaginates(t *testing.T) {
	var calls int
	client := retryablehttp.NewClient()
	client.RetryMax = 0
	client.Logger = nil
	client.HTTPClient = &http.Client{Transport: roundTripperFunc(func(req *http.Request) (*http.Response, error) {
		calls++
		if calls == 1 {
			return activityJSONResponse(http.StatusOK, `{"items":[{"id":31,"name":"Completed"}],"count":2,"next_href":"https://api.gb.intelliflo.net/v2/activities/outcomes?skip=500"}`), nil
		}
		return activityJSONResponse(http.StatusOK, `{"items":[{"id":32,"name":"Cancelled"}],"count":2}`), nil
	})}

	service := activityServiceWithClient(client)
	got, err := service.GetAllOutcomes()
	if err != nil {
		t.Fatalf("GetAllOutcomes() error = %v", err)
	}
	if len(got.Items) != 2 || got.Items[1].ID != 32 {
		t.Fatalf("items = %#v", got.Items)
	}
}

func TestGetAllCategoriesRejectsCollectionFailures(t *testing.T) {
	tests := []struct {
		name      string
		responses []*http.Response
		wantError string
	}{
		{
			name:      "non-200",
			responses: []*http.Response{activityJSONResponse(http.StatusTeapot, `{"message":"upstream failed"}`)},
			wantError: "status code: 418",
		},
		{
			name:      "malformed JSON",
			responses: []*http.Response{activityJSONResponse(http.StatusOK, `{`)},
			wantError: "error parsing body",
		},
		{
			name:      "missing cursor skip",
			responses: []*http.Response{activityJSONResponse(http.StatusOK, `{"items":[],"next_href":"https://api.gb.intelliflo.net/v2/activities/categories"}`)},
			wantError: "invalid next cursor",
		},
		{
			name: "repeated cursor",
			responses: []*http.Response{
				activityJSONResponse(http.StatusOK, `{"items":[],"next_href":"https://api.gb.intelliflo.net/v2/activities/categories?skip=500"}`),
				activityJSONResponse(http.StatusOK, `{"items":[],"next_href":"https://api.gb.intelliflo.net/v2/activities/categories?skip=500"}`),
			},
			wantError: "repeated next cursor",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			call := 0
			client := retryablehttp.NewClient()
			client.RetryMax = 0
			client.Logger = nil
			client.HTTPClient = &http.Client{Transport: roundTripperFunc(func(*http.Request) (*http.Response, error) {
				response := test.responses[call]
				call++
				return response, nil
			})}

			service := activityServiceWithClient(client)
			_, err := service.GetAllCategories()
			if err == nil || !strings.Contains(err.Error(), test.wantError) {
				t.Fatalf("error = %v, want containing %q", err, test.wantError)
			}
		})
	}
}

func activityServiceWithClient(client *retryablehttp.Client) ActivityService {
	var creds credentials.TenantCredentials
	creds.SetClient(client)
	creds.SetAccessToken("token")
	creds.SetApiKey("api-key")
	return ActivityService{Credentials: &creds}
}

func activityJSONResponse(status int, body string) *http.Response {
	return &http.Response{
		StatusCode: status,
		Body:       io.NopCloser(bytes.NewBufferString(body)),
		Header:     make(http.Header),
	}
}
