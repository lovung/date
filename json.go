package date

import (
	"errors"
	"time"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// UnmarshalJSON to parse the JSON
func (d *Date) UnmarshalJSON(bs []byte) error {
	var s string
	err := json.Unmarshal(bs, &s)
	if err != nil {
		return err
	}
	if s == "" {
		*d = NewZero()
		return nil
	}
	t, err := time.ParseInLocation(RFC3339Date, s, time.UTC)
	if err != nil {
		return err
	}
	*d = New(t)
	return nil
}

// MarshalJSON marshal to the JSON
func (d Date) MarshalJSON() ([]byte, error) {
	if d.IsZero() {
		return json.Marshal(nil)
	}
	var t = d.ToTime()
	if y := t.Year(); y < 0 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return nil, errors.New("date.MarshalJSON: year outside of range [0,9999]")
	}
	return json.Marshal(t.Format(RFC3339Date))
}
