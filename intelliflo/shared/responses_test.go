package shared

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestHandleCustomResponseCodePreservesRequestRejected(t *testing.T) {
	response := &http.Response{
		StatusCode: http.StatusBadRequest,
		Body:       io.NopCloser(strings.NewReader(`[{"field":"Address.Country.Code"}]`)),
	}

	_, err := HandleCustomResponseCode(response, http.StatusOK)
	if !errors.Is(err, ErrRequestRejected) {
		t.Fatalf("error = %v, want ErrRequestRejected", err)
	}
}
