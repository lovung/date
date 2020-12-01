# Date Type Support in Golang
https://github.com/lovung/date/workflows/go/badge.svg
[![codecov](https://codecov.io/gh/lovung/date/branch/main/graph/badge.svg?token=BlpWq5Bmcl)](https://codecov.io/gh/lovung/date)

- [x] Date type in Golang
- [x] Unit testing
- [x] NullDate support
- [x] Some methods to working with Date type

    --
    import "."


## Usage

```go
const RFC3339Date = "2006-01-02"
```
RFC3339Date format

#### func  RemoveTime

```go
func RemoveTime(t time.Time) time.Time
```
RemoveTime removes the time values of time.Time at UTC Ex:
2020-01-01T15:30:30Z07:00 -> 2020-01-01T00:00:00Z

#### type Date

```go
type Date time.Time
```

Date is the non-nullable type for Date only. Support UTC timezone only

#### func  MustParse

```go
func MustParse(s string) Date
```
MustParse creates a new Date from the RFC3339 Date - "2006-01-02" Panic if wrong
format

#### func  New

```go
func New(t time.Time) Date
```
New creates a new Date

#### func  NewDate

```go
func NewDate(year int, month time.Month, day int) Date
```
NewDate from year, month and day

#### func  NewFromStr

```go
func NewFromStr(s string) (Date, error)
```
NewFromStr creates a new Date from the RFC3339 Date - "2006-01-02"

#### func  NewZero

```go
func NewZero() Date
```
NewZero creates the new zero (null) Date

#### func (Date) AddDate

```go
func (d Date) AddDate(year, month, day int) Date
```
AddDate returns the Date corresponding to adding the given number of years,
months, and days to d. For example, AddDate(-1, 2, 3) applied to January 1, 2011
returns March 4, 2010.

AddDate normalizes its result in the same way that Date does, so, for example,
adding one month to October 31 yields December 1, the normalized form for
November 31.

#### func (Date) AddMonths

```go
func (d Date) AddMonths(n int) Date
```
AddMonths return add the number of months. If the day at this month > maxDay
(like 30 for Feb), it will be the last day of this month

#### func (Date) After

```go
func (d Date) After(ref Date) bool
```
After compares date if it's after or not

#### func (Date) Before

```go
func (d Date) Before(ref Date) bool
```
Before compares date if it's before or not

#### func (Date) Day

```go
func (d Date) Day() int
```
Day returns the day of the NullDate

#### func (Date) DaysOfMonth

```go
func (d Date) DaysOfMonth() int
```
DaysOfMonth returns total days in the month of the Date

#### func (Date) DiffMonths

```go
func (d Date) DiffMonths(ref Date) int
```
DiffMonths to calculate the diff of months If date.After(ref) -> Return negative
If date.Before(ref) -> Return positive Id date.Equal(ref) -> Return 1
a.DiffMonths(b) may != b.DiffMonths(a) because of the last day of month

#### func (Date) Equal

```go
func (d Date) Equal(ref Date) bool
```
Equal to compare with another

#### func (Date) GormDataType

```go
func (d Date) GormDataType() string
```
GormDataType gorm common data type

#### func (Date) IsZero

```go
func (d Date) IsZero() bool
```
IsZero check if it is the zero value

#### func (Date) MarshalJSON

```go
func (d Date) MarshalJSON() ([]byte, error)
```
MarshalJSON marshal to the JSON

#### func (Date) Month

```go
func (d Date) Month() time.Month
```
Month returns the month of the Date

#### func (*Date) Scan

```go
func (d *Date) Scan(value interface{}) (err error)
```
Scan implements the Scanner interface

#### func (Date) String

```go
func (d Date) String() string
```
Strong converts the Date to string as RFC3339 format

#### func (Date) Sub

```go
func (d Date) Sub(ref Date) time.Duration
```
Sub return duration date - ref

#### func (Date) ToTime

```go
func (d Date) ToTime() time.Time
```
ToTime converts the Date type to time.Time type

#### func (*Date) UnmarshalJSON

```go
func (d *Date) UnmarshalJSON(bs []byte) error
```
UnmarshalJSON to parse the JSON

#### func (Date) Value

```go
func (d Date) Value() (driver.Value, error)
```
Value implements the driver Valuer interface

#### func (Date) YMD

```go
func (d Date) YMD() (year int, month time.Month, day int)
```
YMD returns year, month and day of the Date

#### func (Date) Year

```go
func (d Date) Year() int
```
Year returns the year of the Date

