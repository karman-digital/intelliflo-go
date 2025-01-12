package sharedmodels

import (
	"fmt"
	"time"
)

type IntellifloDateTime time.Time

func (x *IntellifloDateTime) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		*x = IntellifloDateTime(time.Time{})
		return nil
	}
	strTime := string(b[1 : len(b)-1])
	t, err := time.Parse("2006-01-02T15:04:05", strTime)
	if err != nil {
		t, err = time.Parse("2006-01-02T15:04:05Z", strTime)
		if err != nil {
			t, err = time.Parse("2006-01-02T15:04:05.000Z", strTime)
			if err != nil {
				return err
			}
		}
	}
	*x = IntellifloDateTime(t)
	return nil
}

type AccessToken string

func (a AccessToken) String() string {
	return string(a)
}

func (a *AccessToken) Set(s string) error {
	*a = AccessToken(s)
	return nil
}

type APIKey string

func (a APIKey) String() string {
	return string(a)
}

type ExpiresAt time.Time

func (e ExpiresAt) Time() time.Time {
	return time.Time(e)
}

func (e *ExpiresAt) Set(t time.Time) error {
	*e = ExpiresAt(t)
	return nil
}

type TenantId int

func (t TenantId) Int() int {
	return int(t)
}

func (t TenantId) String() string {
	return fmt.Sprintf("%d", t.Int())
}

func (t *TenantId) Set(i int) error {
	*t = TenantId(i)
	return nil
}

type ClientSecret string

func (c ClientSecret) String() string {
	return string(c)
}

func (c *ClientSecret) Set(s string) {
	*c = ClientSecret(s)
}

type ClientId string

func (c ClientId) String() string {
	return string(c)
}

func (c *ClientId) Set(s string) {
	*c = ClientId(s)
}
