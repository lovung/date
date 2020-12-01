# nulldate

## Usage

#### type NullDate

```go
type NullDate struct {
	Date  date.Date
	Valid bool
}
```

NullDate is the nullable type for Date only. Support UTC timezone only Null if
valid is true

#### func  MustParse

```go
func MustParse(s string) NullDate
```
MustParse creates a new Date from the RFC3339 Date - "2006-01-02" Panic if wrong
format

#### func  New

```go
func New(t time.Time, valid bool) NullDate
```
New creates a new Date

#### func  NewDate

```go
func NewDate(year int, month time.Month, day int) NullDate
```
NewDate from year, month and day

#### func  NewFrom

```go
func NewFrom(t time.Time) NullDate
```
NewFrom creates a new Date that will be valid

#### func  NewFromPtr

```go
func NewFromPtr(t *time.Time) NullDate
```
NewFromPtr creates a Date that be null if t is nil

#### func  NewFromStr

```go
func NewFromStr(s string) (NullDate, error)
```
NewFromStr creates a new Date from the RFC3339 Date - "2006-01-02"

#### func  NewZero

```go
func NewZero() NullDate
```
NewZero creates the new zero (null) Date

#### func (NullDate) GormDataType

```go
func (d NullDate) GormDataType() string
```
GormDataType gorm common data type

#### func (NullDate) IsZero

```go
func (d NullDate) IsZero() bool
```
IsZero returns true for invalid Date's, for omitempty support

#### func (NullDate) MarshalJSON

```go
func (d NullDate) MarshalJSON() ([]byte, error)
```
MarshalJSON marshal to the JSON

#### func (*NullDate) Scan

```go
func (d *NullDate) Scan(value interface{}) (err error)
```
Scan implements the Scanner interface

#### func (*NullDate) SetValid

```go
func (d *NullDate) SetValid(t time.Time)
```
SetValid changes this Date's value and also sets it to be non-null

#### func (NullDate) String

```go
func (d NullDate) String() string
```
Strong converts the Date to string as RFC3339 format It will be empty string if
the Date is null

#### func (NullDate) ToTime

```go
func (d NullDate) ToTime() time.Time
```
ToTime returns a time.Time to this Date's value

#### func (*NullDate) UnmarshalJSON

```go
func (d *NullDate) UnmarshalJSON(bs []byte) error
```
UnmarshalJSON to parse the JSON

#### func (NullDate) Value

```go
func (d NullDate) Value() (driver.Value, error)
```
Value implements the driver Valuer interface
