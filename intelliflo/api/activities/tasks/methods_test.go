package tasks

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"testing"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/karman-digital/intelliflo-go/intelliflo/api/credentials"
	sharedmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/shared"
	taskmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/tasks"
)

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestCreateTaskUsesWriteOnlyRequestAndDecodesResponse(t *testing.T) {
	request := taskmodels.TaskCreateRequest{
		Subject:       "Client review",
		Description:   "Staging proof",
		ActivityType:  taskmodels.TaskType{ID: 293522, Category: sharedmodels.IOSubObject{ID: 28380}},
		Priority:      &sharedmodels.IOSubObject{ID: 10070},
		RelatedTo:     []taskmodels.RelatedEntity{{ID: 41589479, Type: "Client"}},
		AssignedTo:    taskmodels.TaskAssignment{User: taskmodels.TaskUser{ID: 678652}},
		DueAt:         "2026-09-15T12:00:00Z",
		CreatedByUser: taskmodels.TaskUser{ID: 678652},
	}
	client := testClient(func(req *http.Request) *http.Response {
		if req.Method != http.MethodPost || req.URL.Path != "/v2/activities/tasks" {
			t.Fatalf("request = %s %s", req.Method, req.URL.Path)
		}
		body, err := io.ReadAll(req.Body)
		if err != nil {
			t.Fatal(err)
		}
		var got taskmodels.TaskCreateRequest
		if err := json.Unmarshal(body, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got, request) {
			t.Fatalf("body = %#v, want %#v", got, request)
		}
		var fields map[string]any
		if err := json.Unmarshal(body, &fields); err != nil {
			t.Fatal(err)
		}
		if _, present := fields["reference"]; present {
			t.Fatal("reference must not be sent on task creation")
		}
		return taskJSONResponse(http.StatusCreated, `{"id":77,"subject":"Client review","reference":"IOT77"}`)
	})
	got, err := taskServiceWithClient(client).CreateTask(request)
	if err != nil || got.ID != 77 {
		t.Fatalf("task = %#v, error = %v", got, err)
	}
}

func TestCreateTaskNoteUsesWriteOnlyRequestAndDecodesResponse(t *testing.T) {
	request := taskmodels.TaskNoteCreateRequest{Notes: "proof marker", ShowOnClientPortal: false, CreatedByUser: taskmodels.TaskUser{ID: 678652}}
	client := testClient(func(req *http.Request) *http.Response {
		if req.Method != http.MethodPost || req.URL.Path != "/v2/activities/tasks/77/notes" {
			t.Fatalf("request = %s %s", req.Method, req.URL.Path)
		}
		var got taskmodels.TaskNoteCreateRequest
		if err := json.NewDecoder(req.Body).Decode(&got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got, request) {
			t.Fatalf("body = %#v, want %#v", got, request)
		}
		return taskJSONResponse(http.StatusCreated, `{"id":88,"notes":"proof marker","showOnClientPortal":false}`)
	})
	got, err := taskServiceWithClient(client).CreateTaskNote(77, request)
	if err != nil || got.ID != 88 {
		t.Fatalf("note = %#v, error = %v", got, err)
	}
}

func TestDeleteTaskRequiresNoContent(t *testing.T) {
	client := testClient(func(req *http.Request) *http.Response {
		if req.Method != http.MethodDelete || req.URL.Path != "/v2/activities/tasks/77" {
			t.Fatalf("request = %s %s", req.Method, req.URL.Path)
		}
		return taskJSONResponse(http.StatusBadGateway, `{"message":"upstream"}`)
	})
	if err := taskServiceWithClient(client).DeleteTask(77); err == nil {
		t.Fatal("expected delete status error")
	}
}

func TestTaskReadsRejectMalformedJSON(t *testing.T) {
	client := testClient(func(*http.Request) *http.Response {
		return taskJSONResponse(http.StatusOK, `{`)
	})
	if _, err := taskServiceWithClient(client).GetTask(77); err == nil {
		t.Fatal("expected decode error")
	}
}

func TestGetTaskUsesActivitiesTasksPath(t *testing.T) {
	client := retryablehttp.NewClient()
	client.RetryMax = 0
	client.Logger = nil
	client.HTTPClient = &http.Client{Transport: roundTripperFunc(func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("method = %s, want GET", req.Method)
		}
		if req.URL.Path != "/v2/activities/tasks/42" {
			t.Fatalf("path = %s, want /v2/activities/tasks/42", req.URL.Path)
		}
		return taskJSONResponse(http.StatusOK, `{"id":42,"subject":"Review"}`), nil
	})}

	service := taskServiceWithClient(client)
	if _, err := service.GetTask(42); err != nil {
		t.Fatalf("GetTask() error = %v", err)
	}
}

func TestGetTaskNoteUsesExactPath(t *testing.T) {
	client := testClient(func(req *http.Request) *http.Response {
		if req.Method != http.MethodGet || req.URL.Path != "/v2/activities/tasks/42/notes/84" {
			t.Fatalf("request = %s %s", req.Method, req.URL.Path)
		}
		return taskJSONResponse(http.StatusOK, `{"id":84,"notes":"Proof note","showOnClientPortal":false}`)
	})

	note, err := taskServiceWithClient(client).GetTaskNote(42, 84)
	if err != nil || note.ID != 84 {
		t.Fatalf("note = %#v, error = %v", note, err)
	}
}

func taskServiceWithClient(client *retryablehttp.Client) *TaskService {
	var creds credentials.TenantCredentials
	creds.SetClient(client)
	creds.SetAccessToken("token")
	creds.SetApiKey("api-key")
	return &TaskService{Credentials: &creds}
}

func taskJSONResponse(status int, body string) *http.Response {
	return &http.Response{
		StatusCode: status,
		Header:     make(http.Header),
		Body:       io.NopCloser(bytes.NewBufferString(body)),
	}
}

func testClient(respond func(*http.Request) *http.Response) *retryablehttp.Client {
	client := retryablehttp.NewClient()
	client.RetryMax = 0
	client.Logger = nil
	client.HTTPClient = &http.Client{Transport: roundTripperFunc(func(req *http.Request) (*http.Response, error) {
		if respond == nil {
			panic("unexpected request")
		}
		return respond(req), nil
	})}
	return client
}
