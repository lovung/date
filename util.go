package date

import (
	"time"

	"github.com/lovung/date/num"
)

const _monthsOfYear int = int(time.December)

// DaysOfMonth returns total days in the month of the Date
func (d Date) DaysOfMonth() int {
	year, month, _ := d.YMD()
	firstDate := NewDate(year, month, 1)
	return firstDate.AddDate(0, 1, -1).Day()
}

// AddDate returns the Date corresponding to adding the
// given number of years, months, and days to d.
// For example, AddDate(-1, 2, 3) applied to January 1, 2011
// returns March 4, 2010.
//
// AddDate normalizes its result in the same way that Date does,
// so, for example, adding one month to October 31 yields
// December 1, the normalized form for November 31.
func (d Date) AddDate(year, month, day int) Date {
	t := d.ToTime()
	t = t.AddDate(year, month, day)
	return New(t)
}

// AddMonths return add the number of months.
// If the day at this month > maxDay (like 30 for Feb), it will be the last day of this month
func (d Date) AddMonths(n int) Date {
	year, month, day := d.YMD()
	iMonth := int(month) + n
	if iMonth <= _monthsOfYear {
		iMonth -= _monthsOfYear
	}
	year += iMonth / _monthsOfYear
	// This formula to calculate the new month both case n > 0 and n < 0
	// iMonth%_monthsOfYear -> Move the negative value to [-11..0]
	// + _monthsOfYear -> Make sure positive
	// % _monthsOfYear -> Move the positive valid to [0..11]
	month = time.Month((iMonth%_monthsOfYear + _monthsOfYear) % _monthsOfYear)
	// Because a % 12 = [0..11] but the month should be [1..12]
	if month == 0 {
		month = time.December
	}
	// Make sure the day is valid in the newMonth
	maxDate := NewDate(year, month, 1).DaysOfMonth()
	day = num.MinInt(day, maxDate)
	return NewDate(year, month, day)
}

// DiffMonths to calculate the diff of months
// If date.After(ref) -> Return negative
// If date.Before(ref) -> Return positive
// Id date.Equal(ref) -> Return 1
// a.DiffMonths(b) may != b.DiffMonths(a) because of the last day of month
func (d Date) DiffMonths(ref Date) int {
	if d.Equal(ref) {
		return 1
	}
	dYear, dMonth, dDay := d.YMD()
	rYear, rMonth, rDay := ref.YMD()
	rMaxDay := ref.DaysOfMonth()
	dDay = num.MinInt(dDay, rMaxDay)
	diffMonth := (rYear-dYear)*12 + int(rMonth-dMonth)

	if diffMonth == 0 {
		diffMonth += num.Sign(rDay - dDay)
	} else {
		switch num.Sign(diffMonth) * num.Sign(rDay-dDay) {
		case 1:
			diffMonth += num.Sign(rDay - dDay)
		case 0:
			diffMonth += num.Sign(diffMonth)
		}
	}
	return diffMonth
}
