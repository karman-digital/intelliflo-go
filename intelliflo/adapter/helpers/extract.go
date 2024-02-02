package intelliflohelpers

import (
	"fmt"
	"net/url"
	"strconv"
)

func ExtractSkipValueFromIntellifloResponse(nextHref string) (int, error) {
	u, err := url.Parse(nextHref)
	if err != nil {
		return 0, err
	}
	queryParams := u.Query()
	skipValueStr := queryParams.Get("skip")
	if skipValueStr == "" {
		return 0, fmt.Errorf("skip parameter not found")
	}
	skipValue, err := strconv.Atoi(skipValueStr)
	if err != nil {
		return 0, fmt.Errorf("failed to convert skip value to integer: %v", err)
	}
	return skipValue, nil
}
