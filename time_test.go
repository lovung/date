package date

import (
	"testing"
	"time"
)

func TestDate_Before(t *testing.T) {
	d1 := Date(time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC))
	d2 := Date(time.Date(2020, 01, 21, 0, 0, 0, 0, time.UTC))
	d0 := Date(time.Time{})
	type args struct {
		ref Date
	}
	tests := []struct {
		name string
		d    Date
		args args
		want bool
	}{
		{
			d: d1,
			args: args{
				ref: d2,
			},
			want: true,
		},
		{
			d: d2,
			args: args{
				ref: d1,
			},
			want: false,
		},
		{
			d: d0,
			args: args{
				ref: d1,
			},
			want: true,
		},
		{
			d: d1,
			args: args{
				ref: d0,
			},
			want: false,
		},
		{
			d: d1,
			args: args{
				ref: d1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.d
			if got := d.Before(tt.args.ref); got != tt.want {
				t.Errorf("Date.Before() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDate_After(t *testing.T) {
	d1 := Date(time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC))
	d2 := Date(time.Date(2020, 01, 21, 0, 0, 0, 0, time.UTC))
	d0 := Date(time.Time{})
	type args struct {
		ref Date
	}
	tests := []struct {
		name string
		d    Date
		args args
		want bool
	}{
		{
			d: d1,
			args: args{
				ref: d2,
			},
			want: false,
		},
		{
			d: d2,
			args: args{
				ref: d1,
			},
			want: true,
		},
		{
			d: d0,
			args: args{
				ref: d1,
			},
			want: false,
		},
		{
			d: d1,
			args: args{
				ref: d0,
			},
			want: true,
		},
		{
			d: d1,
			args: args{
				ref: d1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.d
			if got := d.After(tt.args.ref); got != tt.want {
				t.Errorf("Date.After() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDate_YMD(t *testing.T) {
	d1 := Date(time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC))
	d0 := Date(time.Time{})
	tests := []struct {
		name      string
		d         Date
		wantYear  int
		wantMonth time.Month
		wantDay   int
	}{
		{
			d:         d1,
			wantYear:  2020,
			wantMonth: 1,
			wantDay:   1,
		},
		{
			d:         d0,
			wantYear:  1,
			wantMonth: 1,
			wantDay:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.d
			gotYear, gotMonth, gotDay := d.YMD()
			if gotYear != tt.wantYear {
				t.Errorf("Date.YMD() gotYear = %v, want %v", gotYear, tt.wantYear)
			}
			if gotMonth != tt.wantMonth {
				t.Errorf("Date.YMD() gotMonth = %v, want %v", gotMonth, tt.wantMonth)
			}
			if gotDay != tt.wantDay {
				t.Errorf("Date.YMD() gotDay = %v, want %v", gotDay, tt.wantDay)
			}
		})
	}
}

func TestDate_Year(t *testing.T) {
	d1 := Date(time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC))
	d0 := Date(time.Time{})
	tests := []struct {
		name string
		d    Date
		want int
	}{
		{
			d:    d1,
			want: 2020,
		},
		{
			d:    d0,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.d
			if got := d.Year(); got != tt.want {
				t.Errorf("Date.Year() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDate_Month(t *testing.T) {
	d1 := Date(time.Date(2020, 12, 01, 0, 0, 0, 0, time.UTC))
	d0 := Date(time.Time{})
	tests := []struct {
		name string
		d    Date
		want time.Month
	}{
		{
			d:    d1,
			want: 12,
		},
		{
			d:    d0,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.d
			if got := d.Month(); got != tt.want {
				t.Errorf("Date.Month() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDate_Day(t *testing.T) {
	d1 := Date(time.Date(2020, 12, 31, 0, 0, 0, 0, time.UTC))
	d0 := Date(time.Time{})
	tests := []struct {
		name string
		d    Date
		want int
	}{
		{
			d:    d1,
			want: 31,
		},
		{
			d:    d0,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.d
			if got := d.Day(); got != tt.want {
				t.Errorf("Date.Day() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDate_Sub(t *testing.T) {
	d1 := MustParse("2020-01-01")
	d2 := MustParse("2020-01-02")
	d3 := MustParse("2020-01-03")
	type args struct {
		ref Date
	}
	tests := []struct {
		name string
		d    Date
		args args
		want time.Duration
	}{
		{
			d: d2,
			args: args{
				ref: d1,
			},
			want: time.Hour * 24,
		},
		{
			d: d2,
			args: args{
				ref: d3,
			},
			want: -time.Hour * 24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.d
			if got := d.Sub(tt.args.ref); got != tt.want {
				t.Errorf("Date.Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDate_Equal(t *testing.T) {
	d1 := MustParse("2020-01-01")
	d2 := MustParse("2020-01-02")
	type args struct {
		ref Date
	}
	tests := []struct {
		name string
		d    Date
		args args
		want bool
	}{
		{
			d: d1,
			args: args{
				ref: d2,
			},
			want: false,
		},
		{
			d: d1,
			args: args{
				ref: d1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.d
			if got := d.Equal(tt.args.ref); got != tt.want {
				t.Errorf("Date.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}
