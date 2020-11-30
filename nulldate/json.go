package nulldate

import (
	"bytes"
	"errors"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/lovung/date"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// UnmarshalJSON to parse the JSON
func (d *NullDate) UnmarshalJSON(bs []byte) error {
	var s string
	err := json.Unmarshal(bs, &s)
	if err != nil {
		return err
	}
	if s == "" {
		*d = NewZero()
		return nil
	}
	t, err := time.ParseInLocation(date.RFC3339Date, s, time.UTC)
	if err != nil {
		return err
	}
	*d = NewFrom(t)
	return nil
}

// MarshalJSON marshal to the JSON
func (d NullDate) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	if !d.Valid {
		b.WriteString(`null`)
		return b.Bytes(), nil
	}
	var t = d.ToTime()
	if y := t.Year(); y < 0 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return nil, errors.New("date.MarshalJSON: year outside of range [0,9999]")
	}
	b.WriteString(`"`)
	b.WriteString(t.Format(date.RFC3339Date))
	b.WriteString(`"`)
	return b.Bytes(), nil
}
