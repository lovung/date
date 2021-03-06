package date

import (
	"database/sql"
	"database/sql/driver"
	"reflect"
)

// Scan implements the Scanner interface
func (d *Date) Scan(value interface{}) (err error) {
	nullTime := &sql.NullTime{}
	err = nullTime.Scan(value)
	if err != nil {
		return err
	}
	if reflect.TypeOf(value) == nil {
		*d = NewZero()
	} else {
		*d = New(nullTime.Time)
	}
	return nil
}

// Value implements the driver Valuer interface
func (d Date) Value() (driver.Value, error) {
	t := d.ToTime()
	if t.IsZero() {
		return nil, nil
	}
	return t, nil
}

// GormDataType gorm common data type
func (d Date) GormDataType() string {
	return "date"
}
