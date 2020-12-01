package date

import (
	"reflect"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want Date
	}{
		{
			args: args{
				t: time.Date(2020, 01, 01, 15, 30, 30, 123, time.Now().Location()),
			},
			want: Date(time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDate(t *testing.T) {
	type args struct {
		year  int
		month time.Month
		day   int
	}
	tests := []struct {
		name string
		args args
		want Date
	}{
		{
			args: args{
				year:  2020,
				month: 01,
				day:   01,
			},
			want: Date(time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDate(tt.args.year, tt.args.month, tt.args.day); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewZero(t *testing.T) {
	tests := []struct {
		name string
		want Date
	}{
		{
			want: Date(time.Time{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewZero(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFromStr(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    Date
		wantErr bool
	}{
		{
			args: args{
				s: "2020-01-01",
			},
			want:    Date(time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC)),
			wantErr: false,
		},
		{
			args: args{
				s: "2020-13-01",
			},
			want:    Date(time.Time{}),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFromStr(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFromStr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustParse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want Date
	}{
		{
			args: args{
				s: "2020-01-01",
			},
			want: Date(time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MustParse(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MustParse() = %v, want %v", got, tt.want)
			}
		})
	}

	shouldPanic(t, func() { MustParse("2020-13-01") })
}

func TestDate_ToTime(t *testing.T) {
	tests := []struct {
		name string
		d    Date
		want time.Time
	}{
		{
			d:    Date(time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC)),
			want: time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.d
			if got := d.ToTime(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Date.ToTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDate_IsZero(t *testing.T) {
	tests := []struct {
		name string
		d    Date
		want bool
	}{
		{
			d:    Date(time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC)),
			want: false,
		},
		{
			d:    Date(time.Time{}),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.d
			if got := d.IsZero(); got != tt.want {
				t.Errorf("Date.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func shouldPanic(t *testing.T, f func()) {
	defer func() { recover() }()
	f()
	t.Errorf("should have panicked")
}
