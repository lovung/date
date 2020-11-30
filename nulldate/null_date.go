package nulldate

import (
	"time"

	"github.com/lovung/date"
)

// NullDate is the nullable type for Date only.
// Support UTC timezone only
// Null if valid is true
type NullDate struct {
	Date  date.Date
	Valid bool
}

// New creates a new Date
func New(t time.Time, valid bool) NullDate {
	return NullDate{
		Date:  date.New(t),
		Valid: valid,
	}
}

// NewFrom creates a new Date that will be valid
func NewFrom(t time.Time) NullDate {
	return NullDate{
		Date:  date.New(t),
		Valid: true,
	}
}

// NewDate from year, month and day
func NewDate(year int, month time.Month, day int) NullDate {
	t := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	return NewFrom(t)
}

// NewZero creates the new zero (null) Date
func NewZero() NullDate {
	return NullDate{
		Date:  date.NewZero(),
		Valid: false,
	}
}

// NewFromPtr creates a Date that be null if t is nil
func NewFromPtr(t *time.Time) NullDate {
	if t == nil {
		return NewZero()
	}
	return NullDate{
		Date:  date.New(*t),
		Valid: true,
	}
}

// NewFromStr creates a new Date from the RFC3339 Date - "2006-01-02"
func NewFromStr(s string) (NullDate, error) {
	t, err := time.Parse(date.RFC3339Date, s)
	if err != nil {
		return NewZero(), err
	}
	return NewFrom(t), nil
}

// MustParse creates a new Date from the RFC3339 Date - "2006-01-02"
// Panic if wrong format
func MustParse(s string) NullDate {
	t, err := time.Parse(date.RFC3339Date, s)
	if err != nil {
		panic(err)
	}
	return NewFrom(t)
}

// ToTime returns a time.Time to this Date's value
func (d NullDate) ToTime() time.Time {
	if !d.Valid {
		return time.Time{}
	}
	return d.Date.ToTime()
}

// SetValid changes this Date's value and also sets it to be non-null
func (d *NullDate) SetValid(t time.Time) {
	d.Date = date.New(t)
	d.Valid = true
}

// IsZero returns true for invalid Date's, for omitempty support
func (d NullDate) IsZero() bool {
	return !d.Valid
}

func removeTime(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}
