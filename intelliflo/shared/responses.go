package shared

import (
	"fmt"
	"io"
	"net/http"
)

func HandleCustomResponseCode(resp *http.Response, code int) (rawBody []byte, err error) {
	rawBody, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading body: %s", err)
	}
	if resp.StatusCode != code {
		if resp.StatusCode == 404 {
			return rawBody, ErrResourceNotFound
		}
		return rawBody, fmt.Errorf("error returned by endpoint: %+v", resp.StatusCode)
	}
	return rawBody, nil
}
