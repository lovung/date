package nulldate

import (
	"database/sql"
	"database/sql/driver"
)

// Scan implements the Scanner interface
func (d *NullDate) Scan(value interface{}) (err error) {
	nullTime := &sql.NullTime{}
	err = nullTime.Scan(value)
	if err != nil {
		return err
	}
	*d = New(nullTime.Time, nullTime.Valid)
	return nil
}

// Value implements the driver Valuer interface
func (d NullDate) Value() (driver.Value, error) {
	t := d.ToTime()
	if t.IsZero() {
		return nil, nil
	}
	return t, nil
}

// GormDataType gorm common data type
func (d NullDate) GormDataType() string {
	return "date"
}
