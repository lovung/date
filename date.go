package date

import "time"

// Date is the nullable type for Date only.
// Support UTC timezone only
// Null if valid is true
type Date struct {
	Date  time.Time
	Valid bool
}

// _RFC3339Date format
const _RFC3339Date = "2006-01-02"

// New creates a new Date
func New(t time.Time, valid bool) Date {
	return Date{
		Date:  removeTime(t),
		Valid: valid,
	}
}

// NewFrom creates a new Date that will be valid
func NewFrom(t time.Time) Date {
	return Date{
		Date:  removeTime(t),
		Valid: true,
	}
}

// NewDate from year, month and day
func NewDate(year int, month time.Month, day int) Date {
	t := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	return NewFrom(t)
}

// NewZero creates the new zero (null) Date
func NewZero() Date {
	return Date{
		Date:  time.Time{},
		Valid: false,
	}
}

// NewFromPtr creates a Date that be null if t is nil
func NewFromPtr(t *time.Time) Date {
	if t == nil {
		return NewZero()
	}
	return Date{
		Date:  removeTime(*t),
		Valid: true,
	}
}

// NewFromStr creates a new Date from the RFC3339 Date - "2006-01-02"
func NewFromStr(s string) (Date, error) {
	t, err := time.Parse(_RFC3339Date, s)
	if err != nil {
		return NewZero(), err
	}
	return NewFrom(t), nil
}

// MustParse creates a new Date from the RFC3339 Date - "2006-01-02"
// Panic if wrong format
func MustParse(s string) Date {
	t, err := time.Parse(_RFC3339Date, s)
	if err != nil {
		panic(err)
	}
	return NewFrom(t)
}

// Ptr returns a pointer to this Date's value, or a nil pointer if this Date is null
func (d Date) Ptr() *time.Time {
	if !d.Valid {
		return nil
	}
	return &d.Date
}

// GetValue returns the value to this Date's value,
// Panic if null
func (d Date) GetValue() time.Time {
	if !d.Valid {
		panic("nil date")
	}
	return d.Date
}

// SetValid changes this Date's value and also sets it to be non-null
func (d *Date) SetValid(t time.Time) {
	d.Date = removeTime(t)
	d.Valid = true
}

// IsZero returns true for invalid Date's, for omitempty support
func (d Date) IsZero() bool {
	return !d.Valid
}

func removeTime(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}
