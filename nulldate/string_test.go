package nulldate

import (
	"testing"
	"time"

	"github.com/lovung/date"
)

func TestDate_String(t *testing.T) {
	type fields struct {
		Date  date.Date
		Valid bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			fields: fields{
				Date:  date.New(time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)),
				Valid: true,
			},
			want: "2020-01-01",
		},
		{
			fields: fields{
				Date:  date.NewZero(),
				Valid: false,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NullDate{
				Date:  tt.fields.Date,
				Valid: tt.fields.Valid,
			}
			if got := d.String(); got != tt.want {
				t.Errorf("Date.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
