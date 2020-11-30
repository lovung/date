package date

import "time"

// Date is the non-nullable type for Date only.
// Support UTC timezone only
type Date time.Time

// New creates a new Date
func New(t time.Time) Date {
	return Date(RemoveTime(t))
}

// NewDate from year, month and day
func NewDate(year int, month time.Month, day int) Date {
	t := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	return New(t)
}

// NewZero creates the new zero (null) Date
func NewZero() Date {
	return Date(time.Time{})
}

// NewFromStr creates a new Date from the RFC3339 Date - "2006-01-02"
func NewFromStr(s string) (Date, error) {
	t, err := time.Parse(RFC3339Date, s)
	if err != nil {
		return NewZero(), err
	}
	return New(t), nil
}

// MustParse creates a new Date from the RFC3339 Date - "2006-01-02"
// Panic if wrong format
func MustParse(s string) Date {
	d, err := NewFromStr(s)
	if err != nil {
		panic(err)
	}
	return d
}

// ToTime converts the Date type to time.Time type
func (d Date) ToTime() time.Time {
	return time.Time(d)
}

// IsZero check if it is the zero value
func (d Date) IsZero() bool {
	return d.ToTime().IsZero()
}

// RemoveTime removes the time values of time.Time at UTC
// Ex: 2020-01-01T15:30:30Z07:00 -> 2020-01-01T00:00:00Z
func RemoveTime(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}
