package validate_test

import (
	"testing"

	iovalidate "github.com/karman-digital/intelliflo/intelliflo/adapter/validate"
	"github.com/stretchr/testify/assert"
)

type validateIntellifloSignatureTest struct {
	name      string
	secret    string
	signature string
	body      []byte
	wantErr   bool
}

var validateIntellifloSignatureTests = []validateIntellifloSignatureTest{
	{
		name:      "Pass: Valid signature",
		secret:    "sdiBw7x57CZqCtGKyUKuFR5QWBdfApW",
		signature: "7bf291dccae4cb5798d79175bf59fb9601bc81e9",
		body:      []byte(`{"test": {"test1": "test","test2": 56000,"test3": true}}`),
		wantErr:   false,
	},
	{
		name:      "Fail: Invalid signature - incorrect secret",
		secret:    "fehuafhIHAfisFWEFaEFAD",
		signature: "7bf291dccae4cb5798d79175bf59fb9601bc81e9",
		body:      []byte(`{"test": {"test1": "test","test2": 56000,"test3": true}}`),
		wantErr:   true,
	},
	{
		name:      "Fail: Invalid signature - body incorrect",
		secret:    "sdiBw7x57CZqCtGKyUKuFR5QWBdfApW",
		signature: "7bf291dccae4cb5798d79175bf59fb9601bc81e9",
		body:      []byte(`{"test": {"test1": "tet","test2": 56000,"test3": true}}`),
		wantErr:   true,
	},
	{
		name:      "Fail: Body json malformed",
		secret:    "sdiBw7x57CZqCtGKyUKuFR5QWBdfApW",
		signature: "7bf291dccae4cb5798d79175bf59fb9601bc81e9",
		body:      []byte(`{"test": {"test1": "tet","test2": 56000,"test3": true}`),
		wantErr:   true,
	},
}

func TestValidateWebhookSignature(t *testing.T) {
	for _, test := range validateIntellifloSignatureTests {
		t.Run(test.name, func(t *testing.T) {
			err := iovalidate.ValidateWebhookSignature(test.secret, test.signature, test.body)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
