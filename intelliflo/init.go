package intelliflo

import "github.com/hashicorp/go-retryablehttp"

func InitIntellifloAPI() IntellfloAPI {
	return &Credentials{
		Client: retryablehttp.NewClient(),
	}
}
