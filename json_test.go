package date

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestDate_UnmarshalJSON(t *testing.T) {
	d0 := Date(time.Time{})
	d1 := MustParse("2020-01-01")
	type args struct {
		bs []byte
	}
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
				bs: []byte{},
			},
			want:    &d0,
			wantErr: true,
		},
		{
			d: &d0,
			args: args{
				bs: []byte(""),
			},
			want:    &d0,
			wantErr: true,
		},
		{
			d: &d0,
			args: args{
				bs: []byte("1234"),
			},
			want:    &d0,
			wantErr: true,
		},
		{
			d: &d0,
			args: args{
				bs: []byte("\"\""),
			},
			want:    &d0,
			wantErr: false,
		},
		{
			d: &d0,
			args: args{
				bs: []byte("\"12345-13-01\""),
			},
			want:    &d0,
			wantErr: true,
		},
		{
			d: &d0,
			args: args{
				bs: []byte("\"2020-01-01\""),
			},
			want:    &d1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.d
			if err := d.UnmarshalJSON(tt.args.bs); (err != nil) != tt.wantErr {
				t.Errorf("Date.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(d, tt.want); diff != "" {
				t.Errorf("Date.UnmarshalJSON() diff = %v", diff)
			}
		})
	}
}

func TestDate_MarshalJSON(t *testing.T) {
	d0 := Date(time.Time{})
	d1 := MustParse("2020-01-01")
	tests := []struct {
		name    string
		d       Date
		want    []byte
		wantErr bool
	}{
		{
			d:       d0,
			want:    []byte("null"),
			wantErr: false,
		},
		{
			d:       d1,
			want:    []byte("\"2020-01-01\""),
			wantErr: false,
		},
		{
			d:       New(time.Date(12345, time.January, 1, 0, 0, 0, 0, time.UTC)),
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.d
			got, err := d.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Date.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Date.MarshalJSON() diff = %v", diff)
			}
		})
	}
}
