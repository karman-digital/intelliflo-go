package helpers_test

import (
	"fmt"
	"testing"

	intelliflohelpers "github.com/karman-digital/intelliflo/intelliflo/helpers"
	"github.com/stretchr/testify/assert"
)

func TestExtractSkipValueFromIntellifloResponse(t *testing.T) {
	var extractSkipValueFromIntellifloResponseTests = []struct {
		name           string
		nextHref       string
		expectedOutput int
		expectedError  error
	}{
		{
			name:           "Pass",
			nextHref:       "https://api.intelliflo.com/v2.0/clients?skip=100",
			expectedOutput: 100,
			expectedError:  nil,
		},
		{
			name:           "Fail - no skip value",
			nextHref:       "https://api.intelliflo.com/v2.0/clients",
			expectedOutput: 0,
			expectedError:  fmt.Errorf("skip parameter not found"),
		},
		{
			name:           "Fail - invalid skip value",
			nextHref:       "https://api.intelliflo.com/v2.0/clients?skip=invalid",
			expectedOutput: 0,
			expectedError:  fmt.Errorf("failed to convert skip value to integer: strconv.Atoi: parsing \"invalid\": invalid syntax"),
		},
	}
	for _, test := range extractSkipValueFromIntellifloResponseTests {
		t.Run(test.name, func(t *testing.T) {
			output, err := intelliflohelpers.ExtractSkipValueFromIntellifloResponse(test.nextHref)
			assert.Equal(t, test.expectedOutput, output)
			if test.expectedError != nil {
				assert.Error(t, err)
			}
		})
	}
}
