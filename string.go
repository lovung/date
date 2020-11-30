package date

// Strong converts the Date to string as RFC3339 format
func (d Date) String() string {
	return d.ToTime().Format(RFC3339Date)
}
