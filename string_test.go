package date

import (
	"testing"
	"time"
)

func TestDate_String(t *testing.T) {
	type fields struct {
		Date  time.Time
		Valid bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			fields: fields{
				Date:  time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
				Valid: true,
			},
			want: "2020-01-01",
		},
		{
			fields: fields{
				Date:  time.Time{},
				Valid: false,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Date{
				Date:  tt.fields.Date,
				Valid: tt.fields.Valid,
			}
			if got := d.String(); got != tt.want {
				t.Errorf("Date.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
