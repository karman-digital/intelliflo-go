package credentials

import (
	"fmt"
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
	intelliflomodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models"
)

func (c *TenantCredentials) SendRequest(method, path string, body []byte, opts ...intelliflomodels.GetOptions) (*http.Response, error) {
	req, err := retryablehttp.NewRequest(method, "https://api.gb.intelliflo.net/v2"+path, body)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header["x-api-key"] = []string{c.ApiKey().String()}
	req.Header["authorization"] = []string{fmt.Sprintf("Bearer %s", c.AccessToken())}
	if len(opts) > 0 {
		q := req.URL.Query()
		if opts[0].Skip != 0 {
			q.Add("skip", fmt.Sprintf("%d", opts[0].Skip))
		}
		if opts[0].Top != 0 {
			q.Add("top", fmt.Sprintf("%d", opts[0].Top))
		} else {
			q.Add("top", "500")
		}
		if opts[0].Filter != "" {
			q.Add("filter", opts[0].Filter)
		}
		req.URL.RawQuery = q.Encode()
	} else {
		q := req.URL.Query()
		q.Add("top", "500")
		req.URL.RawQuery = q.Encode()
	}
	resp, err := c.Client().Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	return resp, nil
}
