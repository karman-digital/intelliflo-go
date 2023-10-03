package intelliflo

import "github.com/hashicorp/go-retryablehttp"

func InitIntellifloAPI() IntellfloAPI {
	return &credentials{
		Client: retryablehttp.NewClient(),
	}
}
