package date

import (
	"testing"
	"time"
)

func TestDate_String(t *testing.T) {
	d1 := Date(time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC))
	d0 := Date(time.Time{})
	tests := []struct {
		name string
		d    Date
		want string
	}{
		{
			d:    d1,
			want: "2020-01-01",
		},
		{
			d:    d0,
			want: "0001-01-01",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.d
			if got := d.String(); got != tt.want {
				t.Errorf("Date.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
