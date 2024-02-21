package intelliflomodels

import "time"

type intellifloDateTime time.Time

func (x *intellifloDateTime) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		*x = intellifloDateTime(time.Time{})
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
	*x = intellifloDateTime(t)
	return nil
}
