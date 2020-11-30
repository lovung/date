package date

import (
	"database/sql/driver"
	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestDate_Scan(t *testing.T) {
	type args struct {
		value interface{}
	}

	d1 := Date(time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC))
	d0 := Date(time.Time{})
	tests := []struct {
		name    string
		d       *Date
		args    args
		want    *Date
		wantErr bool
	}{
		{
			d: &d0,
			args: args{
				value: time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC),
			},
			want:    &d1,
			wantErr: false,
		},
		{
			d: &d0,
			args: args{
				value: "2020-01-01",
			},
			want:    &d0,
			wantErr: true,
		},
		{
			d: &d0,
			args: args{
				value: nil,
			},
			want:    &d0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.d
			if err := d.Scan(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Date.Scan() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(d, tt.want); diff != "" {
				t.Errorf("Date.Scan() diff = %v", diff)
			}
		})
	}
}

func TestDate_Value(t *testing.T) {
	d1 := Date(time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC))
	d0 := Date(time.Time{})
	tests := []struct {
		name    string
		d       Date
		want    driver.Value
		wantErr bool
	}{
		{
			d:       d1,
			want:    time.Time(d1),
			wantErr: false,
		},
		{
			d:       d0,
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.d
			got, err := d.Value()
			if (err != nil) != tt.wantErr {
				t.Errorf("Date.Value() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Date.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDate_GormDataType(t *testing.T) {
	d1 := Date(time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC))
	tests := []struct {
		name string
		d    Date
		want string
	}{
		{
			d:    d1,
			want: "date",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.d
			if got := d.GormDataType(); got != tt.want {
				t.Errorf("Date.GormDataType() = %v, want %v", got, tt.want)
			}
		})
	}
}
