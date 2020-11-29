package date

// Strong converts the Date to string as RFC3339 format
// It will be empty string if the Date is null
func (d Date) String() string {
	if !d.Valid {
		return ""
	}
	return d.Ptr().Format(_RFC3339Date)
}
