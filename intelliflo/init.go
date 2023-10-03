package intelliflo

import "github.com/hashicorp/go-retryablehttp"

func InitCredentials() IntellfloAPI {
	return &Credentials{
		Client: retryablehttp.NewClient(),
	}
}
