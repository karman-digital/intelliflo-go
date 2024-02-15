package intelliflo

import (
	"log"
	"os"

	"github.com/hashicorp/go-retryablehttp"
)

func InitIntellifloAPI() IntellfloAPI {
	client := retryablehttp.NewClient()
	client.Logger = log.New(os.Stderr, "", log.LstdFlags)
	return &credentials{
		client: client,
	}
}
