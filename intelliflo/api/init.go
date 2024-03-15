package intelliflo

import (
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
)

func InitIntellifloAPI() IntellifloAPI {
	client := retryablehttp.NewClient()
	client.Logger = nil
	client.ErrorHandler = func(resp *http.Response, err error, numTries int) (*http.Response, error) {
		return resp, nil
	}
	return &credentials{
		client: client,
	}
}
