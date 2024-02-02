package intelliflo

import "github.com/hashicorp/go-retryablehttp"

func InitIntellifloAPI() IntellfloAPI {
	client := retryablehttp.NewClient()
	client.Logger = nil
	return &credentials{
		client: client,
	}
}
