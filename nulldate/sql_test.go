package nulldate

import (
	"database/sql/driver"
	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/lovung/date"
)

func TestNullDate_Scan(t *testing.T) {
	d1 := New(time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC), true)
	d0 := date.New(time.Time{})
	d2 := New(time.Time{}, false)
	type fields struct {
		Date  date.Date
		Valid bool
	}
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *NullDate
		wantErr bool
	}{
		{
			fields: fields{
				Date:  d0,
				Valid: false,
			},
			args: args{
				value: time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC),
			},
			want:    &d1,
			wantErr: false,
		},
		{
			fields: fields{
				Date:  d0,
				Valid: false,
			},
			args: args{
				value: "2020-01-01",
			},
			want:    &d2,
			wantErr: true,
		},
		{
			fields: fields{
				Date:  d0,
				Valid: false,
			},
			args: args{
				value: nil,
			},
			want:    &d2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &NullDate{
				Date:  tt.fields.Date,
				Valid: tt.fields.Valid,
			}
			if err := d.Scan(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("NullDate.Scan() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(d, tt.want); diff != "" {
				t.Errorf("Date.Scan() diff = %v", diff)
			}
		})
	}
}

func TestNullDate_Value(t *testing.T) {
	d1 := date.New(time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC))
	d3 := New(time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC), true)
	d0 := date.New(time.Time{})
	type fields struct {
		Date  date.Date
		Valid bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    driver.Value
		wantErr bool
	}{
		{
			fields: fields{
				Date:  d1,
				Valid: true,
			},
			want:    d3.ToTime(),
			wantErr: false,
		},
		{
			fields: fields{
				Date:  d0,
				Valid: true,
			},
			want:    nil,
			wantErr: false,
		},
		{
			fields: fields{
				Date:  d0,
				Valid: false,
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NullDate{
				Date:  tt.fields.Date,
				Valid: tt.fields.Valid,
			}
			got, err := d.Value()
			if (err != nil) != tt.wantErr {
				t.Errorf("NullDate.Value() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NullDate.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullDate_GormDataType(t *testing.T) {
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
				Date:  date.Date{},
				Valid: false,
			},
			want: "date",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NullDate{
				Date:  tt.fields.Date,
				Valid: tt.fields.Valid,
			}
			if got := d.GormDataType(); got != tt.want {
				t.Errorf("NullDate.GormDataType() = %v, want %v", got, tt.want)
			}
		})
	}
}
