package nulldate

import (
	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/lovung/date"
)

func TestNew(t *testing.T) {
	type args struct {
		t     time.Time
		valid bool
	}
	tests := []struct {
		name string
		args args
		want NullDate
	}{
		{
			args: args{
				t:     time.Date(2020, 01, 01, 15, 30, 30, 123, time.Now().Location()),
				valid: false,
			},
			want: NullDate{
				Date:  date.New(time.Date(2020, 01, 01, 15, 30, 30, 123, time.Now().Location())),
				Valid: false,
			},
		},
		{
			args: args{
				t:     time.Date(2020, 01, 01, 15, 30, 30, 123, time.Now().Location()),
				valid: true,
			},
			want: NullDate{
				Date:  date.New(time.Date(2020, 01, 01, 15, 30, 30, 123, time.Now().Location())),
				Valid: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.t, tt.args.valid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFrom(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want NullDate
	}{
		{
			args: args{
				t: time.Date(2020, 01, 01, 15, 30, 30, 123, time.Now().Location()),
			},
			want: NullDate{
				Date:  date.New(time.Date(2020, 01, 01, 15, 30, 30, 123, time.Now().Location())),
				Valid: true,
			},
		},
		{
			args: args{
				t: time.Time{},
			},
			want: NullDate{
				Date:  date.NewZero(),
				Valid: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFrom(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFrom() = %v, want %v", got, tt.want)
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
		want NullDate
	}{
		{
			args: args{
				year:  2020,
				month: 01,
				day:   01,
			},
			want: NullDate{date.New(time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC)), true},
		},
		{
			args: args{
				year:  2020,
				month: 02,
				day:   30,
			},
			want: NullDate{date.New(time.Date(2020, 03, 01, 0, 0, 0, 0, time.UTC)), true},
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
		want NullDate
	}{
		{
			want: NullDate{
				Date:  date.NewZero(),
				Valid: false,
			},
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

func TestNewFromPtr(t *testing.T) {
	t0 := time.Time{}
	t1 := time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC)
	type args struct {
		t *time.Time
	}
	tests := []struct {
		name string
		args args
		want NullDate
	}{
		{
			args: args{
				t: nil,
			},
			want: NullDate{
				Date:  date.Date{},
				Valid: false,
			},
		},
		{
			args: args{
				t: &t0,
			},
			want: NullDate{
				Date:  date.New(t0),
				Valid: true,
			},
		},
		{
			args: args{
				t: &t1,
			},
			want: NullDate{
				Date:  date.New(t1),
				Valid: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromPtr(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromPtr() = %v, want %v", got, tt.want)
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
		want    NullDate
		wantErr bool
	}{
		{
			args: args{
				s: "2020-01-01",
			},
			want:    NullDate{date.Date(time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC)), true},
			wantErr: false,
		},
		{
			args: args{
				s: "2020-13-01",
			},
			want:    NullDate{date.Date(time.Time{}), false},
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
		want NullDate
	}{
		{
			args: args{
				s: "2020-01-01",
			},
			want: NullDate{date.Date(time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC)), true},
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

func TestNullDate_ToTime(t *testing.T) {
	type fields struct {
		Date  date.Date
		Valid bool
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NullDate{
				Date:  tt.fields.Date,
				Valid: tt.fields.Valid,
			}
			if got := d.ToTime(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NullDate.ToTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullDate_SetValid(t *testing.T) {
	type fields struct {
		Date  date.Date
		Valid bool
	}
	type args struct {
		t time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *NullDate
	}{
		{
			fields: fields{
				Date:  date.Date{},
				Valid: false,
			},
			args: args{
				t: time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC),
			},
			want: &NullDate{
				Date:  date.New(time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC)),
				Valid: true,
			},
		},
		{
			fields: fields{
				Date:  date.Date{},
				Valid: true,
			},
			args: args{
				t: time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC),
			},
			want: &NullDate{
				Date:  date.New(time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC)),
				Valid: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &NullDate{
				Date:  tt.fields.Date,
				Valid: tt.fields.Valid,
			}
			d.SetValid(tt.args.t)
			if diff := cmp.Diff(d, tt.want); diff != "" {
				t.Errorf("Date.Scan() diff = %v", diff)
			}
		})
	}
}

func TestNullDate_IsZero(t *testing.T) {
	type fields struct {
		Date  date.Date
		Valid bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			fields: fields{
				Date:  date.Date{},
				Valid: false,
			},
			want: true,
		},
		{
			fields: fields{
				Date:  date.Date{},
				Valid: true,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NullDate{
				Date:  tt.fields.Date,
				Valid: tt.fields.Valid,
			}
			if got := d.IsZero(); got != tt.want {
				t.Errorf("NullDate.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeTime(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			args: args{
				t: time.Date(2020, 01, 01, 15, 30, 30, 12345, time.UTC),
			},
			want: time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeTime(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func shouldPanic(t *testing.T, f func()) {
	defer func() { recover() }()
	f()
	t.Errorf("should have panicked")
}
