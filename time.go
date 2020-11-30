package date

import "time"

// Before compares date if it's before or not
func (d Date) Before(ref Date) bool {
	return d.GetValue().Before(ref.GetValue())
}

// After compares date if it's after or not
func (d Date) After(ref Date) bool {
	return d.GetValue().After(ref.GetValue())
}

// YMD returns year, month and day of the Date
func (d Date) YMD() (year int, month time.Month, day int) {
	if !d.Valid {
		panic("nil date")
	}
	return d.Date.Date()
}

// Year returns the year of the Date
func (d Date) Year() int {
	y, _, _ := d.YMD()
	return y
}

// Month returns the month of the Date
func (d Date) Month() time.Month {
	_, m, _ := d.YMD()
	return m
}

// Day returns the day of the Date
func (d Date) Day() int {
	_, _, day := d.YMD()
	return day
}

// Sub return duration date - ref
func (d Date) Sub(ref Date) time.Duration {
	return d.GetValue().Sub(ref.GetValue())
}

// Equal to compare with another
func (d Date) Equal(ref Date) bool {
	return d.String() == ref.String()
}
