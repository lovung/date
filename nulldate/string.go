package nulldate

import "github.com/lovung/date"

// Strong converts the Date to string as RFC3339 format
// It will be empty string if the Date is null
func (d NullDate) String() string {
	if !d.Valid {
		return ""
	}
	return d.ToTime().Format(date.RFC3339Date)
}
