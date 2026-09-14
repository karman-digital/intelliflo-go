package activities

import (
	"encoding/json"
	"fmt"
	"net/http"

	sharedmodels "github.com/karman-digital/intelliflo-go/intelliflo/api/models/shared"
	intelliflohelpers "github.com/karman-digital/intelliflo-go/intelliflo/helpers"
	"github.com/karman-digital/intelliflo-go/intelliflo/shared"
)

type activityCollection[T any] struct {
	Href      string `json:"href"`
	FirstHref string `json:"first_href"`
	LastHref  string `json:"last_href"`
	NextHref  string `json:"next_href"`
	PrevHref  string `json:"prev_href"`
	Items     []T    `json:"items"`
	Count     int    `json:"count"`
}

func getAllActivityCollection[T any](service *ActivityService, path string) (activityCollection[T], error) {
	var all activityCollection[T]
	options := sharedmodels.GetOptions{Top: 500}
	seenCursors := map[string]struct{}{}
	seenSkips := map[int]struct{}{0: {}}

	for {
		page, err := getActivityCollectionPage[T](service, path, options)
		if err != nil {
			return activityCollection[T]{}, err
		}
		if all.Href == "" {
			all.Href = page.Href
			all.FirstHref = page.FirstHref
			all.LastHref = page.LastHref
			all.PrevHref = page.PrevHref
			all.Count = page.Count
		}
		all.NextHref = page.NextHref
		all.Items = append(all.Items, page.Items...)

		if page.NextHref == "" {
			return all, nil
		}
		if _, exists := seenCursors[page.NextHref]; exists {
			return activityCollection[T]{}, fmt.Errorf("repeated next cursor: %s", page.NextHref)
		}
		seenCursors[page.NextHref] = struct{}{}

		skip, err := intelliflohelpers.ExtractSkipValueFromIntellifloResponse(page.NextHref)
		if err != nil {
			return activityCollection[T]{}, fmt.Errorf("invalid next cursor: %w", err)
		}
		if _, exists := seenSkips[skip]; exists {
			return activityCollection[T]{}, fmt.Errorf("repeated next cursor skip: %d", skip)
		}
		seenSkips[skip] = struct{}{}
		options.Skip = skip
	}
}

func getActivityCollectionPage[T any](service *ActivityService, path string, options sharedmodels.GetOptions) (activityCollection[T], error) {
	var page activityCollection[T]
	resp, err := service.SendRequest(http.MethodGet, path, nil, options)
	if err != nil {
		return page, fmt.Errorf("error making get request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := shared.HandleCustomResponseCode(resp, http.StatusOK)
	if err != nil {
		return page, fmt.Errorf("error returned by endpoint, status code: %d: %w", resp.StatusCode, err)
	}
	if err := json.Unmarshal(respBody, &page); err != nil {
		return page, fmt.Errorf("error parsing body: %w", err)
	}
	return page, nil
}
